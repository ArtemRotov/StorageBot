package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "you wrote:"+inputMsg.Text)
	c.bot.Send(msg)
}
