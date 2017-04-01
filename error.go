package ampCacheURL

import (
	"fmt"
)

type BadURLError struct {
	rawURL      string
	messageBase string
}

func (err BadURLError) Error() string {
	return fmt.Sprintf(err.messageBase, err.rawURL)
}
