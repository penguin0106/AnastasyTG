package keyboard

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// AdminKeyboard возвращает клавиатуру для администраторов
func AdminKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/manage_users"),
			tgbotapi.NewKeyboardButton("/manage_orders"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/manage_products"),
		),
	)
	return keyboard
}
