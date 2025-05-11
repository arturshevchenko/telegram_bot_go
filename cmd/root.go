package cmd

import (
	"log"
	"os"
	"time"

	"github.com/arturshevchenko/telegram_bot_go/handlers"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gopkg.in/telebot.v4"
)

var rootCmd = &cobra.Command{
	Use:   "telegram-bot",
	Short: "Telegram bot with Cobra and Telebot",
	Run: func(cmd *cobra.Command, args []string) {
		// Load .env file if present
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found or error loading .env file:", err)
		}

		token := os.Getenv("TELE_TOKEN")
		if token == "" {
			log.Fatal("TELE_TOKEN env variable not set")
		}

		pref := telebot.Settings{
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}

		bot, err := telebot.NewBot(pref)
		if err != nil {
			log.Fatal(err)
		}

		handlers.Register(bot)

		log.Println("Bot is running...")
		bot.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
