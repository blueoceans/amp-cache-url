package main

import (
	"flag"
	"fmt"

	"github.com/blueoceans/amp-cache-url"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		var (
			result string
			err    error
		)
		if ampCacheURL.IsCacheURL(arg) {
			result, err = ampCacheURL.GetOriginURL(arg)
		} else {
			result, err = ampCacheURL.GetCacheURL(arg)
		}
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		fmt.Println(result)
	}
}
