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
		c.client.Send(inputMsg.Chat.ID, "bad string")
		return
	}
	strs := strings.Split(res[0], " ")
	err := c.dataAccessObject.Add(inputMsg.From.ID, strs[1], strs[2], strs[3])
	if err != nil {
		log.Printf("error sql insert")
		c.client.Send(inputMsg.Chat.ID, "error sql insert")
		return
	}

	c.client.Send(inputMsg.Chat.ID, "Success!")
}
