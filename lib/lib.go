package main

import (
	"C"

	"github.com/blueoceans/amp-cache-url"
)

func init() {
}

func main() {
}

//export isCacheURL
func isCacheURL(rawurl string) bool {
	return ampCacheURL.IsCacheURL(rawurl)
}

//export getCacheURL
func getCacheURL(rawurl string) (string, error) {
	return ampCacheURL.GetCacheURL(rawurl)
}

//export getOriginURL
func getOriginURL(rawurl string) (string, error) {
	return ampCacheURL.GetOriginURL(rawurl)
}
