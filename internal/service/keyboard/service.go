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
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Добавить"),
				tgbotapi.NewKeyboardButton("Изменить"),
				tgbotapi.NewKeyboardButton("Удалить"),
				tgbotapi.NewKeyboardButton("Помощь"),
			),
		),
		commands: map[string]string{
			"Список":   "list",
			"Добавить": "add",
			"Изменить": "change",
			"Удалить":  "delete",
			"Помощь":   "help",
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
