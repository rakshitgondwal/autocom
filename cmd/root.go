package cmd

import (
	"fmt"
	"os"

	"github.com/rakshitgondwal/autocom/cmd/auth"
	"github.com/rakshitgondwal/autocom/cmd/msg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version string
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "autocom",
	Short: "Autocom is a CLI tool used to generate auto commit messages.",
	Long: `Autocom is an AI based CLI tool which can be used to generate
			`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Hi welcome!")
	},
}

func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(msg.MsgCommand)
	rootCmd.AddCommand(auth.AuthCmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with the name "autcom" (without extensions).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".autocom")

	}

	viper.AutomaticEnv() //read in envirenment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Print("")
	}
}
