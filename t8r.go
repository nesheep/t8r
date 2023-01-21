package t8r

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

func Slice(s string, opts *Options) []string {
	res := make([]string, 0, len(s))
	for _, v := range s {
		res = append(res, string(v))
	}
	return res
}

func Println(s string, opts *Options) {
	o := initOptions(opts)
	sl := Slice(s, o)
	for _, v := range sl {
		fmt.Print(v)
		time.Sleep(time.Second / time.Duration(o.CPS))
	}
	fmt.Println()
}

func PrintFile(filename string, opts *Options) error {
	o := initOptions(opts)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		Println(line, o)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
