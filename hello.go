package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/wxpjimmy/hello/v2/demo"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(demo.Hi("jimmy", "fr"))
}
