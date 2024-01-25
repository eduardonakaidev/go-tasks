package util

import "github.com/google/uuid"

func NewUUID() string {
	id := uuid.New()
	idString := id.String()
	return idString
}