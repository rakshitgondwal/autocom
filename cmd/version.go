package cmd

import (
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Autocom",
	Long:  `This is the version of Autocom`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("autocom version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
