package hello

import (
	"fmt"
)

// Say returns a greeting message for the provided name.
func Say(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
