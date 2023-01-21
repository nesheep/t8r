package t8r

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"unicode"

	"github.com/alecthomas/chroma/v2/quick"
)

type Options struct {
	CPS int // Characters per second
}

var DefaultOpts = &Options{CPS: 50}

func initOptions(opts *Options) *Options {
	if opts == nil {
		return DefaultOpts
	}
	return opts
}

type Typewriter struct {
	Opts *Options
}

func NewTypewriter(opts *Options) Typewriter {
	return Typewriter{Opts: initOptions(opts)}
}

func (w Typewriter) Write(p []byte) (int, error) {
	for _, v := range string(p) {
		fmt.Print(string(v))
		if unicode.IsLetter(v) {
			time.Sleep(time.Second / time.Duration(w.Opts.CPS))
		}
	}
	return len(p), nil
}

func Println(s string, opts *Options) {
	w := NewTypewriter(opts)
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

	w := NewTypewriter(opts)
	if err := quick.Highlight(w, string(b), "go", "terminal256", "monokai"); err != nil {
		return err
	}

	return nil
}
