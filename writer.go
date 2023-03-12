package t8r

import (
	"fmt"
	"io"
	"time"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

type Writer struct {
	w       io.Writer
	Lexer   chroma.Lexer
	Style   *chroma.Style
	Options *Options
}

func NewWriter(w io.Writer, lexer chroma.Lexer, style *chroma.Style, options *Options) Writer {
	l := lexer
	if l == nil {
		l = lexers.Fallback
	}

	s := style
	if s == nil {
		s = styles.Fallback
	}

	return Writer{
		w:       w,
		Lexer:   l,
		Style:   s,
		Options: newOptions(options),
	}
}

func (w Writer) Write(p []byte) (int, error) {
	if w.Options.Highlighted {
		return w.writeHighlighted(p)
	}
	return w.write(p)
}

func (w Writer) write(p []byte) (int, error) {
	for _, v := range string(p) {
		fmt.Fprint(w.w, string(v))
		time.Sleep(time.Second / time.Duration(w.Options.CPS))
	}
	return len(p), nil
}

func (w Writer) writeHighlighted(p []byte) (int, error) {
	iter, err := w.Lexer.Tokenise(nil, string(p))
	if err != nil {
		return 0, err
	}

	l := 1
	newLine := true
	for t := iter(); t != chroma.EOF; t = iter() {
		if w.Options.WithNumber {
			if newLine {
				fmt.Fprintf(w.w, "%6d  ", l)
				l++
			}
			if t.Value == "\n" {
				newLine = true
			} else {
				newLine = false
			}
		}
		w.writeToken(t)
	}

	return len(p), nil
}

func (w Writer) writeToken(t chroma.Token) {
	e := w.Style.Get(t.Type)
	if e.IsZero() {
		w.write([]byte(t.Value))
		return
	}

	out := ""
	if e.Bold == chroma.Yes {
		out += "\033[1m"
	}
	if e.Underline == chroma.Yes {
		out += "\033[4m"
	}
	if e.Italic == chroma.Yes {
		out += "\033[3m"
	}
	if e.Colour.IsSet() {
		out += fmt.Sprintf("\033[38;2;%d;%d;%dm", e.Colour.Red(), e.Colour.Green(), e.Colour.Blue())
	}

	fmt.Fprint(w.w, out)
	w.write([]byte(t.Value))
	fmt.Fprint(w.w, "\033[0m")
}
