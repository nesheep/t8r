package t8r

import (
	"fmt"
	"io"
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

	style := w.style("monokai")
	lxr := w.lexer(s)
	iter, err := lxr.Tokenise(nil, s)
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

func (w Typewriter) lexer(s string) chroma.Lexer {
	var lxr chroma.Lexer
	if w.Lng != "" {
		lxr = lexers.Get(w.Lng)
	} else {
		lxr = lexers.Analyse(s)
	}
	if lxr == nil {
		lxr = lexers.Fallback
	}
	return lxr
}

func (w Typewriter) style(name string) *chroma.Style {
	s := styles.Get(name)
	if s == nil {
		s = styles.Fallback
	}
	return s
}
