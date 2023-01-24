package t8r

type Options struct {
	CPS         int  // Characters per second
	Highlighted bool // Enable syntax highlighting
}

var DefaultOptions = &Options{
	CPS:         50,
	Highlighted: false,
}

func newOptions(options *Options) *Options {
	if options == nil {
		return DefaultOptions
	}
	return options
}
