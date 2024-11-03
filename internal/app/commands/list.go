package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) List(message *tgbotapi.Message) {
	response := "Here all products:"
	response += "\n"
	response += "\n"
	for _, v := range c.productService.List() {
		response += v.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
