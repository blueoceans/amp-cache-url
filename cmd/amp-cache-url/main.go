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
		if amp_cache_url.IsCacheUrl(arg) {
			result, err = amp_cache_url.GetOriginUrl(arg)
		} else {
			result, err = amp_cache_url.GetCacheUrl(arg)
		}
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		fmt.Println(result)
	}
}
