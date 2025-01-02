package main

import (
	"log"
	"os"
	"github.com/go-telegram-bot-api/telegram-api/v5"
)
func main() {
	bot, err := tgbotapi.NewBotAPI("7927809753:AAGaYqVDw7MY0eYDaaU-0899Zg7k-uPSj6Q")
	if err !=nil{
		log.Panic(err)
	} 
		
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.@mugutdin),

	u := tgbotapi. NewUpdate(0)
	u.Timeout=60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nill {
			log.Printf("[%s] %s", update.Message.From.@mugutdin,update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,update.Message.Text)
			bot.Send(msg)
		}
	}
}
for update := range updates {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.@mugutdin, update.Message.Text)

		var msg tgbotapi.MessageConfig

		switch update .Message.Command() {
		case "start" :
			msg = tgbotapi.NewMessage(update.Message.Chat.ID,"Welcome! I am your bot.")
		case "help":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID,"I can help you with the following commands:\n/start - Start the bot\n/help - Display this help message")
			default:
			    msg = tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
			}

			bot.Send(msg)
		}
	}
