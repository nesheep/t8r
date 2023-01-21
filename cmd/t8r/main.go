package main

import (
	"strings"

	"github.com/nesheep/t8r"
	"github.com/spf13/cobra"
)

func main() {
	echoCmd := &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t8r.Println(strings.Join(args, " "), nil)
		},
	}

	rootCmd := &cobra.Command{Use: "t8r"}
	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}
