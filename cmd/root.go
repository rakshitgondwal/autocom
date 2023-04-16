package cmd

import (
	"fmt"
	"os"

	"autocom/cmd/msg"

	"github.com/spf13/cobra"
)

var (
	version string
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

func init(){
	rootCmd.AddCommand(msg.MsgCommand)
}