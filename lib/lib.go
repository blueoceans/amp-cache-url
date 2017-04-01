package main

import (
	"C"

	"github.com/blueoceans/amp-cache-url"
)

func init() {
}

func main() {
}

//export isCacheUrl
func isCacheUrl(rawurl string) bool {
	return amp_cache_url.IsCacheUrl(rawurl)
}

//export getCacheUrl
func getCacheUrl(rawurl string) (string, error) {
	return amp_cache_url.GetCacheUrl(rawurl)
}

//export getOriginUrl
func getOriginUrl(rawurl string) (string, error) {
	return amp_cache_url.GetOriginUrl(rawurl)
}
