package ampCacheURL

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/net/idna"
)

const (
	cdnAmpprojectOrg    string = "cdn.ampproject.org"
	dotCdnAmpprojectOrg string = ".cdn.ampproject.org"
)

var (
	typeMap = map[string]string{
		"":      "c",
		".htm":  "c",
		".html": "c",

		".gif": "i",
		".jpg": "i",
		".png": "i",
		".svg": "i",

		".otf":   "r",
		".woff":  "r",
		".woff2": "r",
	}
)

func isCacheURL(rawurl string) (bool, *url.URL) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return false, u
	}
	switch {
	case
		u.Host == cdnAmpprojectOrg,
		strings.HasSuffix(u.Host, dotCdnAmpprojectOrg):
		return true, u
	default:
		return false, u
	}
}

func getType(u *url.URL) string {
	ext := filepath.Ext(u.Path)
	if ext == "" {
		ext = filepath.Ext(u.RawQuery)
	}
	return typeMap[ext]
}

// IsCacheURL returns if provide string is AMP-Cache-URL.
func IsCacheURL(rawurl string) bool {
	result, _ := isCacheURL(rawurl)
	return result
}

// GetCacheURL returns AMP-Cache-URL of provide URL or error.
func GetCacheURL(rawurl string) (string, error) {
	result, u := isCacheURL(rawurl)
	if result {
		return "", &BadURLError{rawurl, "must be not AMP-Cache-URL, got %q"}
	} else if u == nil {
		return "", &BadURLError{rawurl, "must be URL, got %q"}
	}

	// https://developers.google.com/amp/cache/overview#amp-cache-url-format
	// Optional 's'
	var scheme string
	switch u.Scheme {
	case "http":
		scheme = ""
	case "https":
		scheme = "/s"
	default:
		return "", &BadURLError{rawurl, "must be HTTP/HTTPS URL, got %q"}
	}

	// Content type
	contentType := getType(u)
	if contentType == "" {
		return "", &BadURLError{rawurl, "un supprted MIME-Type, got %q"}
	}

	// Subdomain name
	// 1. Converting the AMP document domain from IDN (Punycode) to UTF-8.
	subdomain, err := idna.ToUnicode(u.Host)
	if err != nil {
		return "", err
	}
	// 2. Replacing every "-" (dash) with "--"(2 dashes).
	subdomain = strings.Replace(subdomain, "-", "--", -1)
	// 3. Replacing every "." (dot) with a "-" (dash).
	subdomain = strings.Replace(subdomain, ".", "-", -1)
	// 4. Converting back to IDN (Punycode).
	subdomain, err = idna.ToASCII(subdomain)
	if err != nil {
		return "", err
	}

	// URI
	removeScheme := strings.SplitN(rawurl, "://", 2)[1]
	return fmt.Sprintf("https://%s%s/%s%s/%s", subdomain, dotCdnAmpprojectOrg, contentType, scheme, removeScheme), nil
}

// GetOriginURL returns origin URL of provide AMP-Cache-URL or error.
func GetOriginURL(rawurl string) (string, error) {
	result, u := isCacheURL(rawurl)
	if !result {
		return "", &BadURLError{rawurl, "must be AMP-Cache-URL, got %q"}
	} else if u == nil {
		return "", &BadURLError{rawurl, "must be URL, got %q"}
	}

	// https://developers.google.com/amp/cache/overview#amp-cache-url-format
	// URI
	uriSplits := strings.SplitN(rawurl, "/", 5)
	if len(uriSplits) != 5 {
		return "", &BadURLError{rawurl, "must be AMP-Cache-URL, got %q"}
	}
	uri := uriSplits[4]

	// Optional 's'
	var scheme string
	if strings.HasPrefix(uri, "s/") {
		scheme = "https"
		uri = uri[2:]
	} else {
		scheme = "http"
	}

	return fmt.Sprintf("%s://%s", scheme, uri), nil
}
