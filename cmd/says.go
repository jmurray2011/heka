package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Channel struct {
	ChannelName string `mapstructure:"name"`
	Webhook     string `mapstructure:"webhook"`
}

type Config struct {
	Channels []Channel `mapstructure:"channels"`
}

var (
	ChannelArg string
	MessageArg string
)
var config Config

// saysCmd represents the says command
var saysCmd = &cobra.Command{
	Use:   "says",
	Short: "Sends a message to Slack channel with optional attachment and message template support",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&config); err != nil {
			conf_err := fmt.Sprintf("%s", err)
			log.Fatal().Msg(conf_err)
		}
		msg_log := fmt.Sprintf("Sending message '%s' to channel '%s'", MessageArg, ChannelArg)
		log.Debug().Msg(msg_log)
		sendMessage(ChannelArg, MessageArg)
	},
}

func init() {
	rootCmd.AddCommand(saysCmd)

	// define flags
	saysCmd.PersistentFlags().StringVarP(&ChannelArg, "channel", "c", "", "the channel to send a message to")
	saysCmd.MarkPersistentFlagRequired("channel")

	saysCmd.PersistentFlags().StringVarP(&MessageArg, "message", "m", "", "the message to send")
	saysCmd.MarkPersistentFlagRequired("message")
}

func sendMessage(channel, message string) error {
	attachment := slack.Attachment{
		Color:         "good",
		AuthorName:    "heka",
		AuthorSubname: "github.com",
		AuthorLink:    "https://github.com/jmurray2011/heka",
		AuthorIcon:    "https://avatars2.githubusercontent.com/u/652790",
		Text:          message,
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment}}

	for k := range config.Channels {
		if channel == config.Channels[k].ChannelName {
			webhook := config.Channels[k].Webhook
			err := slack.PostWebhook(webhook, &msg)
			if err != nil {
				slack_err := fmt.Sprintf("%s", err)
				log.Fatal().Msg(slack_err)
				return err
			}
			return nil
		}
	}
	err := fmt.Sprintf("channel '%s' is not in the config file", channel)
	log.Fatal().Msg(err)
	return fmt.Errorf(err)
}
