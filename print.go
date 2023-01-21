package t8r

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alecthomas/chroma/v2/lexers"
)

func Println(s string, opts *Options) {
	w := NewTypewriter(os.Stdout, opts)
	fmt.Fprintln(w, s)
}

func PrintFile(filename string, opts *Options) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	lxr := lexers.Match(filename)
	if lxr == nil {
		lxr = lexers.Fallback
	}

	w := NewTypewriter(os.Stdout, opts)
	w.Lng = lxr.Config().Name
	fmt.Fprint(w, string(b))

	return nil
}
