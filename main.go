package main

import (
	"log"
	"os"

	"AnastasyTG/bot"
	"github.com/joho/godotenv"
)

func init() {
	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Получение токена бота из переменной окружения
	botToken := os.Getenv("BOT_TOKEN")

	// Создание нового экземпляра бота
	myBot, err := bot.NewBot(botToken)
	if err != nil {
		log.Fatal("Error creating bot instance:", err)
	}

	// Запуск бота
	if err := myBot.Start(); err != nil {
		log.Fatal("Error starting bot:", err)
	}
}
