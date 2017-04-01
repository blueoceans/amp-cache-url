package main

import (
	"C"

	"github.com/blueoceans/amp-cache-url"
)

func init() {
}

func main() {
}

//export IsCacheURL
func IsCacheURL(rawurl string) bool {
	return ampCacheURL.IsCacheURL(rawurl)
}

//export GetCacheURL
func GetCacheURL(rawurl string) (string, error) {
	return ampCacheURL.GetCacheURL(rawurl)
}

//export GetOriginURL
func GetOriginURL(rawurl string) (string, error) {
	return ampCacheURL.GetOriginURL(rawurl)
}
