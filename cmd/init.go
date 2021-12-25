/*
MIT License

Copyright (c) 2021 Josh Murray

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generates a default config file at $HOME/.heka.json. Will not overwrite existing configs.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// copy example config to $HOME/.heka.yaml if it doesn't exist already
		homedir, err := os.UserHomeDir()
		if err != nil {
			e := fmt.Sprintf("%s", err)
			log.Fatal().Msg(e)
		}
		config_file := fmt.Sprintf("%s/.heka.json", homedir)
		example_config := "lib/.heka.example.json"

		copy(example_config, config_file)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func copy(src, dst string) error {
	sourceFileStat, src_err := os.Stat(src)
	if src_err != nil {
		e := fmt.Sprintf("%s", src_err)
		log.Fatal().Msg(e)
		return src_err
	}

	_, dst_err := os.Stat(dst)
	if dst_err == nil {
		e := fmt.Sprintf("%s already exists, not overwriting", dst)
		log.Fatal().Msg(e)
		return dst_err
	}

	if !sourceFileStat.Mode().IsRegular() {
		e := fmt.Sprintf("%s is not a regular file", src)
		log.Fatal().Msg(e)
		return fmt.Errorf(e)
	}

	source, err := os.Open(src)
	if err != nil {
		e := fmt.Sprintf("%s", err)
		log.Fatal().Msg(e)
		return fmt.Errorf(e)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		e := fmt.Sprintf("%s", err)
		log.Fatal().Msg(e)
		return fmt.Errorf(e)
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		e := fmt.Sprintf("%s", err)
		log.Fatal().Msg(e)
		return fmt.Errorf(e)
	} else {
		i := fmt.Sprintf("config file saved at %s, please update it with the appropriate information", dst)
		log.Info().Msg(i)
		return fmt.Errorf(i)
	}
}
