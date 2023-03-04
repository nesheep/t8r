package t8r

type Options struct {
	CPS         int    // Characters per second
	Style       string // Name of style
	Highlighted bool   // Enable syntax highlighting
}

var DefaultOptions = &Options{
	CPS:         50,
	Style:       "monokai",
	Highlighted: false,
}

func newOptions(options *Options) *Options {
	if options == nil {
		return DefaultOptions
	}
	return options
}
