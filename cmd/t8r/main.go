package main

import (
	"log"
	"strings"

	"github.com/nesheep/t8r"
	"github.com/spf13/cobra"
)

func main() {
	echoCmd := &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Args:  cobra.MinimumNArgs(1),
		Run:   echo,
	}

	catCmd := &cobra.Command{
		Use:   "cat [filenaes]",
		Short: "Print files to the screen",
		Args:  cobra.MinimumNArgs(1),
		Run:   cat,
	}

	rootCmd := &cobra.Command{Use: "t8r"}
	rootCmd.AddCommand(echoCmd, catCmd)
	rootCmd.Execute()
}

func echo(cmd *cobra.Command, args []string) {
	t8r.Println(strings.Join(args, " "), nil)
}

func cat(cmd *cobra.Command, args []string) {
	for _, v := range args {
		err := t8r.PrintFile(v, nil)
		if err != nil {
			log.Print(err)
		}
	}
}
