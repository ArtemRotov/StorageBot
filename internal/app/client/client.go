package client

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	Bot       *tgbotapi.BotAPI
	UpdConfig tgbotapi.UpdateConfig
	token     string
}

func NewClient(token string) *Client {
	c := &Client{
		Bot:       nil,
		UpdConfig: tgbotapi.UpdateConfig{Timeout: 60},
		token:     token,
	}

	c.initialize()

	return c
}

func (c *Client) initialize() {
	var err error
	c.Bot, err = tgbotapi.NewBotAPI(c.token)
	if err != nil {
		log.Panic(err)
	}

	c.Bot.Debug = false

	log.Printf("Authorized on account %s", c.Bot.Self.UserName)
}

func (c *Client) GetUpdatesChan() tgbotapi.UpdatesChannel {
	return c.Bot.GetUpdatesChan(c.UpdConfig)
}

func (c *Client) Send(id int64, msg string) {
	message := tgbotapi.NewMessage(id, msg)
	message.ParseMode = tgbotapi.ModeMarkdown
	c.Bot.Send(message)
}

func (c *Client) SendReplyMarkup(id int64, msg string, reply interface{}) {
	message := tgbotapi.NewMessage(id, msg)
	message.ParseMode = tgbotapi.ModeMarkdown
	message.ReplyMarkup = reply
	c.Bot.Send(message)
}
