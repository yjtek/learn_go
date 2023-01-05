package main

import (
	"fmt"
)

func checkReturn(a string) string {
	var test string
	var test2 string

	switch a {
	case "p2m":
		test = "abc"
		test2 = test
	case "lba":
		test = "efg"
		test2 = test
	default:
		test = "xyz"
		test2 = test
	
	return test2
}

func main() {
	var a string
	fmt.Scan(&a)
	test := checkReturn(a)
	fmt.Println(test)
}
