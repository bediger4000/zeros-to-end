package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	array := prepareCmdLine(os.Args[1:])
	fiddleArray(array)
	fmt.Printf("%v\n", array)
}

func prepareCmdLine(repr []string) []int {
	var ary []int
	for _, str := range repr {
		n, err := strconv.Atoi(str)
		if err == nil {
			ary = append(ary, n)
		}
	}
	return ary
}

func fiddleArray(array []int) {
}
