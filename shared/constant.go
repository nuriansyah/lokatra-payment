package shared

import (
	"strings"
)

func ParseErrorCode(c string) (string, bool) {
	index := strings.Index(c, ":")
	if index == -1 {
		return "", false
	}

	return c[:index], true
}
