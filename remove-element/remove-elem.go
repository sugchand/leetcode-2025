package main

import "fmt"

func removeElement(nums []int, val int) int {
	numLen := len(nums)
	indexList := make([]int, 0)

	for i := 0; i < numLen; i++ {
		if nums[i] != val {
			indexList = append(indexList, nums[i])
		}
	}
	for i, k := range indexList {
		nums[i] = k
	}
	fmt.Print(nums)
	return len(indexList)
}

func main() {
	fmt.Print(removeElement([]int{3, 2, 2, 3}, 4))
}
