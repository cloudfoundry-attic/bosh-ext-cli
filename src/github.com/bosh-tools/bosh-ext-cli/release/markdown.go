package release

import (
	"strings"
)

func markdown(str string) string {
	// tick char is more commonly used than @
	return strings.TrimSpace(strings.Replace(str, "@", "`", -1))
}
