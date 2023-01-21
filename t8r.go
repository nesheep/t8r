package t8r

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

type Options struct {
	CPS         int  // Characters per second
	Highlighted bool // Print syntax highlighted text or not
}

var DefaultOpts = &Options{
	CPS:         50,
	Highlighted: false,
}

func initOptions(opts *Options) *Options {
	if opts == nil {
		return DefaultOpts
	}
	return opts
}

type Typewriter struct {
	Lng  string
	Opts *Options
	w    io.Writer
}

func NewTypewriter(w io.Writer, opts *Options) Typewriter {
	return Typewriter{
		Opts: initOptions(opts),
		w:    w,
	}
}

func (w Typewriter) Write(p []byte) (int, error) {
	if w.Opts.Highlighted {
		return w.writeHighlighted(p)
	}
	return w.write(p)
}

func (w Typewriter) write(p []byte) (int, error) {
	for _, v := range string(p) {
		fmt.Fprint(w.w, string(v))
		time.Sleep(time.Second / time.Duration(w.Opts.CPS))
	}
	return len(p), nil
}

func (w Typewriter) writeHighlighted(p []byte) (int, error) {
	s := string(p)

	var lexer chroma.Lexer
	if w.Lng != "" {
		lexer = lexers.Get(w.Lng)
	} else {
		lexer = lexers.Analyse(s)
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}

	iter, err := lexer.Tokenise(nil, s)
	if err != nil {
		return 0, err
	}

	for token := iter(); token != chroma.EOF; token = iter() {
		entry := style.Get(token.Type)
		if !entry.IsZero() {
			out := ""
			if entry.Bold == chroma.Yes {
				out += "\033[1m"
			}
			if entry.Underline == chroma.Yes {
				out += "\033[4m"
			}
			if entry.Italic == chroma.Yes {
				out += "\033[3m"
			}
			if entry.Colour.IsSet() {
				out += fmt.Sprintf("\033[38;2;%d;%d;%dm", entry.Colour.Red(), entry.Colour.Green(), entry.Colour.Blue())
			}
			fmt.Fprint(w.w, out)
		}
		w.write([]byte(token.Value))
		if !entry.IsZero() {
			fmt.Fprint(w.w, "\033[0m")
		}
	}

	return len(p), nil
}

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

	lexer := lexers.Match(filename)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	w := NewTypewriter(os.Stdout, opts)
	w.Lng = lexer.Config().Name
	fmt.Fprint(w, string(b))

	return nil
}
