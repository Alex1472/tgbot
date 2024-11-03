package main

import (
	"fmt"
	"github.com/Alex1472/tgbot/internal/app/commands"
	"github.com/Alex1472/tgbot/internal/services/product"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	token := os.Getenv("TOKEN")
	fmt.Println(token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 0,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewSerivce()
	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "list":
			commander.List(update.Message)
		default:
			commander.Default(update.Message)
		}
	}
}
