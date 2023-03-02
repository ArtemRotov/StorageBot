package commands

import (
	"database/sql"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sadbard/StorageBot/internal/service/keyboard"
	"github.com/sadbard/StorageBot/internal/storage"
)

type Commander struct {
	bot             *tgbotapi.BotAPI
	keyboardService *keyboard.Service
	recordDB        storage.RecordInterface
}

func NewCommander(bot *tgbotapi.BotAPI, keybServ *keyboard.Service, db *sql.DB) *Commander {
	return &Commander{
		bot:             bot,
		keyboardService: keybServ,
		recordDB:        &storage.RecordDB{DB: db},
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.Message != nil { // If we got a message

		command := update.Message.Command()
		if len(command) == 0 {
			command = c.keyboardService.Сommand(update.Message.Text)
		}

		switch command {
		case "start":
			c.Start(update.Message)
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "add":
			c.Add(update.Message)
		case "change":
			c.Change(update.Message)
		case "delete":
			c.Delete(update.Message)
		default:
			c.Default(update.Message)
		}
	}
}
