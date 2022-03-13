package coreUuid

import "github.com/google/uuid"

// CreateUuid creates a new UUID.
func CreateUuid() string {
	return uuid.New().String()
}
