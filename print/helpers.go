package print

import (
	"fmt"
	"strings"
)

// prepend is a helper function that returns each line of the text
// prepended by the prepend string.
func prepend(prepend, text string) string {
	var s strings.Builder
	for _, t := range strings.SplitAfter("\n", text) {
		s.WriteString(fmt.Sprintf("%s%s\n", prepend, t))
	}
	return s.String()
}
