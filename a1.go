package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	array := prepareCmdLine(os.Args[1:])
	fmt.Printf("Before: %v\n", array)
	zerostoend(array)
	fmt.Printf("After:  %v\n", array)
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

func zerostoend(array []int) {

	var zeroIdx int

	length := len(array)

	for ; zeroIdx < length; zeroIdx++ {
		if array[zeroIdx] != 0 {
			continue
		}
		// zeroIdx is the index of a zero-array-value
		// find non-zero value further along in the array
		for nonZeroIdx := zeroIdx; nonZeroIdx < length; nonZeroIdx++ {
			if array[nonZeroIdx] == 0 {
				continue
			}
			array[zeroIdx] = array[nonZeroIdx]
			array[nonZeroIdx] = 0
			break
		}
	}
}
