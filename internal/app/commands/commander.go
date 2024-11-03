package commands

import (
	"github.com/Alex1472/tgbot/internal/services/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}
