package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sadbard/StorageBot/internal/app/client"
	"github.com/sadbard/StorageBot/internal/service/keyboard"
	"github.com/sadbard/StorageBot/internal/storage/models"
)

type RecordInterface interface {
	All(userId int64) ([]models.Record, error)
	Add(userId int64, label, login, password string) error
}

type ClientInterface interface {
	Send(id int64, msg string)
	SendReplyMarkup(id int64, msg string, reply interface{})
}

type Commander struct {
	client           ClientInterface
	keyboardService  *keyboard.Service
	dataAccessObject RecordInterface
}

func NewCommander(cl *client.Client, keybServ *keyboard.Service, dao RecordInterface) *Commander {
	return &Commander{
		client:           cl,
		keyboardService:  keybServ,
		dataAccessObject: dao,
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
