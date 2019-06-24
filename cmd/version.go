package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of gen-cli",
	Long: "Print the long version number of gen-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0")
	},
}

func Execute() {
	if err := versionCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}