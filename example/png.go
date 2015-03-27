package main

import (
	"fmt"

	"github.com/beefsack/go-dot"
)

func main() {
	fmt.Println(dot.Render([][]bool{
		{false, true, false, true},
		{true, true, false, true},
		{false, true, true, true},
		{false, true, false, true},
	}))
}
