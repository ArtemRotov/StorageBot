package commands

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {

	records, err := c.dataAccessObject.All(inputMsg.From.ID)
	if err != nil {
		log.Fatalln(err)
	}
	strOut := strings.Builder{}
	for _, rec := range records {
		strOut.WriteString(rec.String())
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, strOut.String())
	c.bot.Send(msg)
}
