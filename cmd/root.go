package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autocom",
	Short: "Autocom is a CLI tool used to generate auto commit messages.",
	Long: `Autocom is an AI based CLI tool which can be used to generate
			`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Hi welcome!")
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }