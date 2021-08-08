package core_domain

import "github.com/google/uuid"

func CreateUuid() string {
	return uuid.New().String()
}
