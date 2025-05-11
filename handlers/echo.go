package handlers

import (
	"gopkg.in/telebot.v4"
)

func Register(bot *telebot.Bot) {
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		text := c.Text()
		return c.Send("Ви написали: " + text)
	})
}
