package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var Verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "heka",
	Short: "A simple Slack messaging tool",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initLogging, initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "set output to verbose")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".heka"
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".heka.toml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in. Fail if it is improperly formatted
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Debug().Msgf("%s", err)
	}
}

func initLogging() {
	// set global logging level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if Verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
