package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	num3 := make([]int, 0, m+n)

	j := 0
	i := 0
	for {
		if j == n {
			// we processe entire num2, lets just merge rest of nums1
			num3 = append(num3, nums1[i:m]...)
			break
		}
		if i == m {
			num3 = append(num3, nums2[j:n]...)
			break
		}
		if nums1[i] < nums2[j] {
			num3 = append(num3, nums1[i])
			i++
		} else {
			num3 = append(num3, nums2[j])
			j++
		}
	}
	for k, v := range num3 {
		nums1[k] = v
	}
}
func main() {
	nums1 := []int{1, 2, 3, 0}
	nums2 := []int{2}
	merge(nums1, 3, nums2, 1)
	fmt.Print(nums1)

}
