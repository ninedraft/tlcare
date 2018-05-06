package advicemodel

import (
	"time"

	"github.com/ninedraft/tlcare/pkg/utils/id"
)

type Advice struct {
	ID            string    `storm:"id"`
	Text          string    `storm:"text,index"`
	Likes         uint64    `storm:"likes"`
	LastRetrieved time.Time `storm:"last_retrieved"`
	Retrieved     int       `storm:"retrieved"`
}

func FromText(text string) Advice {
	return Advice{
		ID:   id.ID(),
		Text: text,
	}
}
