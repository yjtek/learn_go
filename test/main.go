package main

import (
	"fmt"

	"github.com/yjtek/learn_go/test/subdir"
)

func main() {
	fmt.Println("test123")

	test := 1

	if test == 1 {
		subdir.SubdirPrint()
	}

}
