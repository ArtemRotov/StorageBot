package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong argument", args)
		return
	}
	pr, err := c.productService.Get(arg)
	if err != nil {
		log.Println("fail product", arg)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, pr.Title)
	c.bot.Send(msg)
}
