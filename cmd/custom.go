package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	versionCmd.AddCommand(customCmd)
}

var customCmd = &cobra.Command{
	Use: "custom",
	Short: "Custom command",
	Long: "This is a custom command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("custom command")
	},
}