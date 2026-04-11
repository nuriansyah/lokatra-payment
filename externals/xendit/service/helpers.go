package service

import (
	"strings"

	"github.com/gofrs/uuid"
)

func parseUUID(value string) uuid.UUID {
	if strings.TrimSpace(value) == "" {
		return uuid.Nil
	}
	id, err := uuid.FromString(value)
	if err != nil {
		return uuid.Nil
	}
	return id
}

func fallbackString(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
