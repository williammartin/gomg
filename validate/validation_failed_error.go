package validate

import (
	"fmt"
	"strings"
)

type ValidationFailedError struct {
	ValidationErrors []string
}

func (e *ValidationFailedError) Error() string {
	var b strings.Builder
	fmt.Fprintln(&b, "validation errors occurred:")
	for _, e := range e.ValidationErrors {
		fmt.Fprintf(&b, " - %s\n", e)
	}

	return b.String()
}
