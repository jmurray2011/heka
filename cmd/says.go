/*
Copyright © 2021 Josh Murray jmurray2011@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"encoding/json"
	"strconv"
	"time"
    "github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Channel struct {
	ChannelName string `mapstructure:"name"`
	Webhook      string `mapstructure:"webhook"`
}

type Config struct {
	Channels []Channel `mapstructure:"channels"`
}

var (
	ChannelArg    string
	TemplateArg   string
	AttachmentArg string
	MessageArg    string
)

// saysCmd represents the says command
var saysCmd = &cobra.Command{
	Use:   "says",
	Short: "Sends a message to Slack channel with optional attachment and message template support",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("Sending message")
		sendMessage(ChannelArg, MessageArg)
	},
}

func init() {
	rootCmd.AddCommand(saysCmd)

	// define flags
	saysCmd.PersistentFlags().StringVarP(&MessageArg, "message", "m", "", "the message to send")
	saysCmd.MarkPersistentFlagRequired("message")

	// each of these have default values specified in the config
	saysCmd.PersistentFlags().StringVarP(&ChannelArg, "channel", "c", "default", "the channel to send a message to")
}

func sendMessage(channel, message string) error {
	if err := viper.ReadInConfig(); err != nil {
		conf_err := fmt.Sprintf("%s", err)
		log.Fatal().Msg(conf_err)
		return err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		conf_err := fmt.Sprintf("%s", err)
		log.Fatal().Msg(conf_err)
		return err
	}

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
		Attachments: []slack.Attachment{attachment},	}

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
		err := fmt.Sprintf("channel '%s' is not in the config file", channel)
		log.Fatal().Msg(err)
	}
	return nil
}

