package util

import (
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GenUUID() string {
	u := uuid.New()
	return u.String()
}
