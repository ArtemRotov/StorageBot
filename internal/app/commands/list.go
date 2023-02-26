package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	outputText := "Here all the products: \n\n"

	products := c.productService.List()
	for _, val := range products {
		outputText += val.Title
		outputText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("NextPage", "some data"),
		),
	)

	c.bot.Send(msg)
}
