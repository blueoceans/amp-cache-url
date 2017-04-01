package ampCacheURL

import (
	"fmt"
)

// BadURLError is error on AMP-Cache-URL.
type BadURLError struct {
	rawURL      string
	messageBase string
}

func (err BadURLError) Error() string {
	return fmt.Sprintf(err.messageBase, err.rawURL)
}
