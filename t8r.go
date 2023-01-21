package t8r

import (
	"fmt"
	"time"
)

type Options struct{}

func Indexes(s string, o *Options) []int {
	res := make([]int, 0, len(s))
	for i := range s {
		res = append(res, i+1)
	}
	return res
}

func Println(s string, o *Options) {
	indexes := Indexes(s, o)
	for i := range indexes {
		start := 0
		if i > 0 {
			start = indexes[i-1]
		}
		fmt.Print(s[start:indexes[i]])
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
}
