package client

import "fmt"

func wrapError(customMessage string, originalError error) error {
	return fmt.Errorf("%s : %v", customMessage, originalError)
}
