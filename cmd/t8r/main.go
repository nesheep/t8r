package main

import (
	"log"
	"strings"

	"github.com/nesheep/t8r"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "t8r"}
var opts t8r.Options

func init() {
	rootCmd.PersistentFlags().IntVarP(&opts.CPS, "cps", "c", t8r.DefaultOpts.CPS, "Characters per second")
}

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

	rootCmd.AddCommand(echoCmd, catCmd)
	rootCmd.Execute()
}

func echo(cmd *cobra.Command, args []string) {
	t8r.Println(strings.Join(args, " "), &opts)
}

func cat(cmd *cobra.Command, args []string) {
	for _, v := range args {
		err := t8r.PrintFile(v, &opts)
		if err != nil {
			log.Print(err)
		}
	}
}
