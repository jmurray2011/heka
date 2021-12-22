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

type Channel struct {
	ChannelName string `mapstructure:"channel-name"`
	Webhook     string `mapstructure:"webhook"`
}

type Config struct {
	ChannelAlias []Channel `mapstructure:"channel"`
}

var ChannelArg string

// saysCmd represents the says command
var saysCmd = &cobra.Command{
	Use:   "says",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			return
		}
		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("says called, %s\n", ChannelArg)
	},
}

func init() {
	rootCmd.AddCommand(saysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	saysCmd.PersistentFlags().StringVarP(&ChannelArg, "channel", "c", "", "the channel to send a message to")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
