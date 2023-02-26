package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Start(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Сервис запущен ✅")
	msg.ReplyMarkup = c.keyboardService.Keyboard
	c.bot.Send(msg)
}
