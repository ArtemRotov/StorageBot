package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Change(inputMsg *tgbotapi.Message) {
	c.client.Send(inputMsg.Chat.ID, "TBD CHANGE")
}
