package contact

import "fmt"

type contactError struct {
	Err    error
	Reason string
}

func (e *contactError) Error() string {
	return fmt.Sprintf("Error: %v, Reason: %v", e.Err, e.Reason)
}
