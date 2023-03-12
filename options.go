package t8r

type Options struct {
	CPS         int  // Characters per second
	WithNumber  bool // Number the output lines, starting at 1
	Highlighted bool // Enable syntax highlighting
}

var DefaultOptions = &Options{
	CPS:         50,
	WithNumber:  false,
	Highlighted: true,
}

func newOptions(options *Options) *Options {
	if options == nil {
		return DefaultOptions
	}
	return options
}
