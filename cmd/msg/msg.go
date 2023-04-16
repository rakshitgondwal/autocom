package msg

import (
	"github.com/spf13/cobra"
)

var MsgCommand = &cobra.Command{
	Use:   "message",
	Aliases: []string{"msg"},
	Short: "Generate a new commit message",
	Long:  `This command will generate a new commit message for
			for the changes you have made to your file/directory`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("new commit message")
	},
}