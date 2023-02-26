package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service struct {
	Keyboard tgbotapi.ReplyKeyboardMarkup
	commands map[string]string
}

func NewService() *Service {
	return &Service{
		Keyboard: tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Список"),
			),
		),
		commands: map[string]string{
			"Список": "list",
		},
	}
}

func (s *Service) Сommand(str string) string {
	val, ok := s.commands[str]
	if !ok {
		return str
	}
	return val
}
