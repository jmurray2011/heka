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
	"io"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generates a default config file at $HOME/.heka.json. Will not overwrite existing configs.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// copy example config to $HOME/.heka.yaml if it doesn't exist already
		config_file := fmt.Sprintf("%s/.heka.json", os.Getenv("HOME"))
		example_config := "lib/.heka.example.json"

		copy(example_config, config_file)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, src_err := os.Stat(src)
	if src_err != nil {
		return 0, src_err
	}

	_, dst_err := os.Stat(dst)
	if dst_err == nil {
		fmt.Printf("%s already exists, not overwriting\n", dst)
		return 0, dst_err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
