package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var (
	cfgFile, subject, natsURL string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cenats",
	Short: "Publish and receive NATS messages in CloudEvents format",
	Long: `cenats is a CLI command that parallels the NATS CLI for publishing/subscribing 
to a NATS message queue in CloudEvents format. This is meant primary for convenient
testing/debugging of CloudEvents services.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cenats.yaml)")
	rootCmd.PersistentFlags().StringVarP(&subject, "subject", "s", "",
		"NATS subject to publish to. In NATS, subjects scope messages into streams or topics. (required)")
	_ = rootCmd.MarkFlagRequired("subject")
	rootCmd.PersistentFlags().StringVarP(&natsURL, "url", "u", "nats://127.0.0.1:4222",
		"URL of the NATS server")
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

		// Search config in home directory with name ".cenats" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cenats")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
