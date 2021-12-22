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

	_ "github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Workspace struct {
	WorkspaceName string    `mapstructure:"workspace-name"`
	OathToken     string    `mapstructure:"oath-token"`
	Channels      []Channel `mapstructure:"channel"`
}

type Channel struct {
	ChannelName string `mapstructure:"channel-name"`
	Webhook     string `mapstructure:"webhook"`
}

type Config struct {
	Workspaces []Workspace `mapstructure:"workspace"`
}

var (
	ChannelArg string
	Template   string
	Attachment string
	Message    string
)

// saysCmd represents the says command
var saysCmd = &cobra.Command{
	Use:   "says",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
			return
		}
		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("says called, %s\n", config.Workspaces[0].Channels)
	},
}

func init() {
	rootCmd.AddCommand(saysCmd)

	// define flags
	saysCmd.PersistentFlags().StringVarP(&ChannelArg, "channel", "c", "", "the channel to send a message to")
	saysCmd.MarkPersistentFlagRequired("channel")

	saysCmd.PersistentFlags().StringVarP(&Message, "message", "m", "", "the message to send")
	saysCmd.MarkPersistentFlagRequired("message")

	saysCmd.PersistentFlags().StringVarP(&Template, "template", "t", "", "set the message template")
	saysCmd.PersistentFlags().StringVarP(&Attachment, "attachment", "a", "", "path to a file to attach to the message")
}
