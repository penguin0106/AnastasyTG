package handlers

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// HandleUserCommand обрабатывает команды от обычных пользователей
func HandleUserCommand(botAPI *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	case "/catalog":
		sendCatalog(botAPI, message.Chat.ID)
	case "/cart":
		sendCart(botAPI, message.Chat.ID)
	case "/order":
		sendOrder(botAPI, message.Chat.ID)
	default:
		log.Printf("Unknown command from user: %s", message.Text)
	}
}

// sendCatalog отправляет каталог товаров пользователю
func sendCatalog(botAPI *tgbotapi.BotAPI, chatID int64) {
	// Здесь ваша реализация отправки каталога товаров
	msg := tgbotapi.NewMessage(chatID, "Это ваш каталог товаров")
	if _, err := botAPI.Send(msg); err != nil {
		log.Println("Ошибка отправки каталога товаров:", err)
	}
}

// sendCart отправляет корзину пользователю
func sendCart(botAPI *tgbotapi.BotAPI, chatID int64) {
	// Здесь ваша реализация отправки корзины
	msg := tgbotapi.NewMessage(chatID, "Это ваша корзина")
	if _, err := botAPI.Send(msg); err != nil {
		log.Println("Ошибка отправки корзины:", err)
	}
}

// sendOrder отправляет форму для оформления заказа
func sendOrder(botAPI *tgbotapi.BotAPI, chatID int64) {
	// Здесь ваша реализация отправки формы заказа
	msg := tgbotapi.NewMessage(chatID, "Заполните форму для оформления заказа")
	if _, err := botAPI.Send(msg); err != nil {
		log.Println("Ошибка отправки формы заказа:", err)
	}
}
