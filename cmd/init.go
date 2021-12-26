package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// setup vars for initCmd
var homedir, _ = os.UserHomeDir()
var short = fmt.Sprintf("Generates a default config file at %s/.heka.toml. Will not overwrite existing configs.", homedir)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: short,
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// copy example config to $HOME/.heka.yaml (or Winddows equivalent) if it doesn't exist already
		config_file := fmt.Sprintf("%s/.heka.toml", homedir)
		example_config := "lib/.heka.example.toml"
		log.Debug().Msgf("Copying %s to %s", example_config, config_file)
		copy(example_config, config_file)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func copy(src, dst string) {
	sourceFileStat, src_err := os.Stat(src)
	if src_err != nil {
		log.Fatal().Msgf("%s doesn't exist: %s", src, src_err)
	}

	_, dst_err := os.Stat(dst)
	if dst_err == nil {
		log.Fatal().Msgf("%s already exists, not overwriting", dst)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Fatal().Msgf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		log.Fatal().Msgf("Couldn't open %s: %s", src, err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		log.Fatal().Msgf("Couldn't create %s: %s", dst, err)
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		log.Fatal().Msgf("Couldn't copy %s to %s: %s", src, dst, err)
	} else {
		log.Info().Msgf("config file saved at %s, please update it with the appropriate information", dst)
	}
}
