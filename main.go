package main

import "fmt"

func main() {
	nums := []int{}
	for i := range 11 {
		nums = append(nums, i)
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
