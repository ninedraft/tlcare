package id

import (
	"fmt"

	"github.com/matoous/go-nanoid"
)

func ID() string {
	id, err := gonanoid.Nanoid()
	if err != nil {
		panic(fmt.Sprintf("[tlcare/pkg/util/id.ID] something completely wrong: %v", err))
	}
	return id
}
