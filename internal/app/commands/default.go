package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) Default(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote: "+message.Text)
	c.bot.Send(msg)
}
