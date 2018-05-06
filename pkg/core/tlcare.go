package core

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Care struct {
	bot *tgbotapi.BotAPI
	db  *storm.DB
}

func NewCare(token string, dbFile string) *Care {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(fmt.Sprintf("[tlcare/pkg/core.NewCare] tgbotapi init: %v", err))
	}
	db, err := storm.Open(dbFile, storm.Codec(json.Codec))
	if err != nil {
		panic(fmt.Sprintf("[tlcare/pkg/core.NewCare] storm db init: %v", err))
	}
	return &Care{
		bot: bot,
		db:  db,
	}
}
