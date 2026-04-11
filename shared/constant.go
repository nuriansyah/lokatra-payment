package shared

import (
	"strings"
)

const (
	ErrorInternalSystem               = "internal system error"
	ErrorInvalidRequest               = "invalid request"
	ErrorNotFound                     = "not found"
	ErrorPQConstrainViolated          = "database constraint violated"
	ErrorPQDuplicateConstrainViolated = "database duplicate constraint violated"
)

func ParseErrorCode(c string) (string, bool) {
	index := strings.Index(c, ":")
	if index == -1 {
		return "", false
	}

	return c[:index], true
}
