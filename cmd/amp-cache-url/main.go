package main

import (
	"flag"
	"fmt"

	"github.com/blueoceans/amp-cache-url"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		if amp_cache_url.IsCacheUrl(arg) {
			if result, err := amp_cache_url.GetOriginUrl(arg); err != nil {
				fmt.Printf("%s\n", err)
			} else {
				fmt.Println(result)
			}
		} else {
			if result, err := amp_cache_url.GetCacheUrl(arg); err != nil {
				fmt.Printf("%s\n", err)
			} else {
				fmt.Println(result)
			}
		}
	}
}
