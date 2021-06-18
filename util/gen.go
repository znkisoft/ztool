package util

import (
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

func genShortId() (string, error) {
	return shortid.Generate()
}

func genUUID() string {
	u := uuid.New()
	return u.String()
}
