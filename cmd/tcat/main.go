package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/nesheep/t8r"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tcat [filenames]",
	Short: "Print files to the screen with typewriter effect",
	Args:  cobra.MinimumNArgs(1),
	Run:   tcat,
}

var stylesCmd = &cobra.Command{
	Use:   "styles [name]",
	Short: "Search names of available styles",
	Args:  cobra.MaximumNArgs(1),
	Run:   searchStyles,
}

var (
	cps   int
	style string
)

func init() {
	styles.Fallback = styles.Get("monokai")

	rootCmd.Flags().IntVarP(&cps, "cps", "c", t8r.DefaultOptions.CPS, "Characters per second")
	rootCmd.Flags().StringVarP(&style, "style", "s", t8r.DefaultOptions.Style, "Name of style")
	rootCmd.AddCommand(stylesCmd)
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
		err := t8r.PrintFile(v, lexers.Match(v), styles.Get(style), options)
		if err != nil {
			log.Print(err)
		}
	}
}

func searchStyles(cmd *cobra.Command, args []string) {
	for _, v := range styles.Names() {
		if len(args) == 0 || strings.HasPrefix(v, args[0]) {
			fmt.Println(v)
		}
	}
}
