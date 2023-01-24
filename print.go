package t8r

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alecthomas/chroma/v2"
)

func Println(s string, lexer chroma.Lexer, style *chroma.Style, options *Options) {
	w := NewWriter(os.Stdout, lexer, style, options)
	fmt.Fprintln(w, s)
}

func PrintlnDefault(s string) {
	Println(s, nil, nil, nil)
}

func PrintFile(filename string, lexer chroma.Lexer, style *chroma.Style, options *Options) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	w := NewWriter(os.Stdout, lexer, style, options)
	fmt.Fprint(w, string(b))

	return nil
}

func PrintFileDefault(filename string) error {
	return PrintFile(filename, nil, nil, nil)
}
