package commands

import (
	"github.com/Alex1472/tgbot/internal/services/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var registeredCommands = map[string]func(*Commander, *tgbotapi.Message){}

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

func (c *Commander) HandleCommand(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	command := update.Message.Command()
	f, ok := registeredCommands[command]
	if ok {
		f(c, update.Message)
	} else {
		c.Default(update.Message)
	}
}
