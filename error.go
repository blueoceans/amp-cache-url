package amp_cache_url

import (
	"fmt"
)

type BadUrlError struct {
	rawUrl      string
	messageBase string
}

func (err BadUrlError) Error() string {
	return fmt.Sprintf(err.messageBase, err.rawUrl)
}
