package bot

import (
	"log"

	"AnastasyTG/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot представляет собой структуру бота
type Bot struct {
	API *tgbotapi.BotAPI
}

// NewBot создает новый экземпляр бота
func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		API: api,
	}, nil
}

// Start запускает бота
func (bot *Bot) Start() error {
	log.Println("Bot started")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.API.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if update.Message.From.UserName == "admin_username" {
				handlers.HandleAdminCommand(bot.API, update.Message)
			} else {
				handlers.HandleUserCommand(bot.API, update.Message)
			}
		}
	}

	return nil
}
