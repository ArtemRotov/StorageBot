package commands

import (
	"log"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Add(inputMsg *tgbotapi.Message) {
	rregexp := `add \w+ \w+ \w+`

	r := regexp.MustCompile(rregexp)
	res := r.FindAllString(inputMsg.Text, -1)
	if len(res) == 0 {
		log.Printf("error regexp")
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "bad string")
		c.bot.Send(msg)
		return
	}
	strs := strings.Split(res[0], " ")
	err := c.dataAccessObject.Add(inputMsg.From.ID, strs[1], strs[2], strs[3])
	if err != nil {
		log.Printf("error sql insert")
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "error sql insert")
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Success!")
	c.bot.Send(msg)
}
