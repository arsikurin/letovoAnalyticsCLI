package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Long: "\033[0;31m          __\n       __/o \\_\n       \\____  \\\n           /   \\\n     __   //\\   \\" +
			"\n  __/o \\-//--\\   \\_/\n  \\____  ___  \\  |\n       ||   \\ |\\ |\n      _||   _||_||\n\n\n\n" +
			"\033[34mHow have you been, man? Drink enough water? \033[0m\n\n",
		Use:     "letovo",
		Short:   "letovo: s.letovo.ru helper",
		Version: "0.7.0-beta",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.letovo.yaml)")

}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			log.Errorln(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".letovo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
