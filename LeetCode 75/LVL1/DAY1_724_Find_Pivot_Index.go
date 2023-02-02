package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	leftSum := 0
	for i, n := range nums {
		if leftSum == sum-leftSum-n {
			return i
		}
		leftSum += n
	}

	return -1
}