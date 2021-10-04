package gen

import (
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

func ShortId() (string, error) {
	return shortid.Generate()
}

func UUID() string {
	return uuid.NewString()
}
