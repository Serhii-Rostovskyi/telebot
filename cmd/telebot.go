/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// telebotCmd represents the telebot command
var telebotCmd = &cobra.Command{
	Use:   "telebot",
	Aliases: []string{"start"},
	Short: "тут може бути Ваша реклама",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("telebot %s started", appVersion)
		tbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("TELE_TOKEN errtext. %s", err)
			return
		}

		tbot.Handle(telebot.OnText, func(m telebot.Context) error {

			log.Print(m.Message().Payload, m.Text())
			payload :=m.Message().Payload

			switch payload {

				case "hello":  
					err = m.Send(fmt.Sprintf("Hello 👋. Im AI %s", appVersion))
				
				case "huilo":  
					err = m.Send("Sam huilo 😡😡😡. ")
				
			}



			return err
		})

		tbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(telebotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// telebotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// telebotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
