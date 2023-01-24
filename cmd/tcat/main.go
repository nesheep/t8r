package main

import (
	"log"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/nesheep/t8r"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tcat [filenaes]",
	Short: "Print files to the screen",
	Args:  cobra.MinimumNArgs(1),
	Run:   tcat,
}

var cps int

func init() {
	rootCmd.PersistentFlags().IntVarP(&cps, "cps", "c", t8r.DefaultOptions.CPS, "Characters per second")
}

func main() {
	rootCmd.Execute()
}

func tcat(cmd *cobra.Command, args []string) {
	options := &t8r.Options{
		CPS:         cps,
		Highlighted: true,
	}
	for _, v := range args {
		err := t8r.PrintFile(v, lexers.Match(v), styles.Get("monokai"), options)
		if err != nil {
			log.Print(err)
		}
	}
}
