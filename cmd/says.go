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
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Workspace struct {
	WorkspaceName  string    `mapstructure:"workspace-name"`
	WorkspaceAlias string    `mapstructure:"workspace-alias"`
	Channels       []Channel `mapstructure:"channel"`
}

type Channel struct {
	ChannelName  string `mapstructure:"channel-name"`
	ChannelAlias string `mapstructure:"channel-alias"`
	Webhook      string `mapstructure:"webhook"`
}

type Config struct {
	Workspaces []Workspace `mapstructure:"workspace"`
}

var (
	WorkspaceArg  string
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
	saysCmd.PersistentFlags().StringVarP(&WorkspaceArg, "workspace", "w", "default", "the workspace the channel lives on")
	saysCmd.PersistentFlags().StringVarP(&TemplateArg, "template", "t", "default", "set the message template")

	// file to attach
	saysCmd.PersistentFlags().StringVarP(&AttachmentArg, "attachment", "a", "", "path to a file to attach to the message")
}

func sendMessage(channel, message string) error {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return err
	}

	attachment := slack.Attachment{
		Color:         "good",
		Fallback:      "You successfully posted by Incoming Webhook URL!",
		AuthorName:    "jmurray2011/heka",
		AuthorSubname: "github.com",
		AuthorLink:    "https://github.com/jmurray2011/heka",
		AuthorIcon:    "https://avatars2.githubusercontent.com/u/652790",
		Text:           message,
		Footer:        "heka message",
		FooterIcon:    "https://platform.slack-edge.com/img/default_application_icon.png",
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},

	}

	for k := range(config.Workspaces) {
		workspace := config.Workspaces[k].WorkspaceAlias
		if workspace == WorkspaceArg{
			fmt.Printf("workspace: %s\n", workspace)
			for k1 := range(config.Workspaces[k].Channels) {
				if channel == config.Workspaces[k].Channels[k1].ChannelAlias {
					fmt.Printf("channel: %s\n", channel)
					if channel == ChannelArg {
						webhook := config.Workspaces[k].Channels[k1].Webhook
						fmt.Printf("webhook: %s\n", webhook)
						err := slack.PostWebhook(webhook, &msg)
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			}
		}
	}

	return nil
}