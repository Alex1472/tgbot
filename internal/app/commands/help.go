package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Help(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "/help - help\n"+"/list - list all products")
	c.bot.Send(msg)
}
