package amp_cache_url

import (
	"testing"
)

var vtests = []struct {
	url         string
	ampCacheURL string
}{
	{"http://example.com/",
		"https://example-com.cdn.ampproject.org/c/example.com/"},
	{"https://example.com/index.html",
		"https://example-com.cdn.ampproject.org/c/s/example.com/index.html"},
	{"https://example.com/i.png",
		"https://example-com.cdn.ampproject.org/i/s/example.com/i.png"},
	{"https://example.com/i.woff",
		"https://example-com.cdn.ampproject.org/r/s/example.com/i.woff"},
}

func TestIsCacheUrl(t *testing.T) {
	for _, vt := range vtests {
		if IsCacheUrl(vt.url) != false {
			t.Errorf("%q, want `false`", vt.url)
		}
		if IsCacheUrl(vt.ampCacheURL) != true {
			t.Errorf("%q, want `true`", vt.ampCacheURL)
		}
	}
}

func TestGetCacheUrl(t *testing.T) {
	for _, vt := range vtests {
		if url, err := GetCacheUrl(vt.url); err != nil {
			t.Errorf("can not get url: %v", err)
		} else if url != vt.ampCacheURL {
			t.Errorf("%q, want %q, got %q", vt.url, vt.ampCacheURL, url)
		}
		if _, err := GetCacheUrl(vt.ampCacheURL); err == nil {
			t.Errorf("%q, want `error`", vt.ampCacheURL)
		}
	}
}

func TestGetOriginUrl(t *testing.T) {
	for _, vt := range vtests {
		if url, err := GetOriginUrl(vt.ampCacheURL); err != nil {
			t.Errorf("can not get url: %v", err)
		} else if url != vt.url {
			t.Errorf("%q, want %q, got %q", vt.ampCacheURL, vt.url, url)
		}
		if _, err := GetOriginUrl(vt.url); err == nil {
			t.Errorf("%q, want `error`", vt.url)
		}
	}
}
