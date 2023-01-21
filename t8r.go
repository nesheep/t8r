package t8r

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Options struct{}

func Slice(s string, o *Options) []string {
	res := make([]string, 0, len(s))
	for _, v := range s {
		res = append(res, string(v))
	}
	return res
}

func Println(s string, o *Options) {
	sl := Slice(s, o)
	for _, v := range sl {
		fmt.Print(v)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

func PrintFile(filename string, o *Options) error {
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
