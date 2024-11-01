package main

import (
	"fmt"
)

func main() {

	var nums = [3]int{1, 2, 3}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	println("=================")

	newNums := []int{0, 1, 2, 3}

	for i := 0; i < len(newNums); i++ {
		fmt.Println(newNums[i])
	}
}
