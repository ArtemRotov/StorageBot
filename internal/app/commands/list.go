package commands

import (
	"database/sql"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sadbard/StorageBot/internal/storage"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	start := time.Now()

	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	recs := storage.RecordDB{}

	duration := time.Since(start).String()

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "TBD LIST")

	c.bot.Send(msg)
}
