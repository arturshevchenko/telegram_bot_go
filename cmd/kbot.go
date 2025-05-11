/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gopkg.in/telebot.v4"
)

// kbotCmd represents the kbot command

func init() {
	rootCmd.AddCommand(kbotCmd)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}
}

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"go"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var TeleToken = os.Getenv("TELE_TOKEN")
		fmt.Printf("token %s \n", TeleToken)
		fmt.Printf("kbot %s started\n", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("TELE_TOKEN env variable must be set and valid. Error: %s", err)
			return
		}
		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload
			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm Kbot %s!", appVersion))
			}

			return err
		})

		kbot.Start()
	},
}
