package keyboard

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// UserKeyboard возвращает клавиатуру для обычных пользователей
func UserKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/catalog"),
			tgbotapi.NewKeyboardButton("/cart"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/order"),
		),
	)
	return keyboard
}
