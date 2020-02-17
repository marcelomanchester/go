package main

import (
	"fmt"
	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Marcelo"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}