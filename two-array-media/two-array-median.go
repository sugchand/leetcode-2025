package main

import (
	"fmt"
	"math"
)

// lets sort the array and create 3rd array and find the middle
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	idx1 := 0
	idx2 := 0
	result_array := []int{}
	if len(nums1) == 0 {
		result_array = nums2
	} else if len(nums2) == 0 {
		result_array = nums1
	} else {
		for {
			curr_val := 0
			if nums1[idx1] < nums2[idx2] {
				curr_val = nums1[idx1]
				idx1++
			} else {
				curr_val = nums2[idx2]
				idx2++
			}
			result_array = append(result_array, curr_val)
			if idx1 == len(nums1) {
				//completed the process of entire one array
				result_array = append(result_array, nums2[idx2:]...)
				break
			} else if idx2 == len(nums2) {
				result_array = append(result_array, nums1[idx1:]...)
				break
			}
		}
	}
	fmt.Print(result_array)
	middle := float64(len(result_array)) / 2.0
	middle_int := int(middle)
	if float64(middle_int) != middle {
		// we found a proper division
		middle_int = int(math.Floor(middle))
		return float64(result_array[middle_int])
	}
	return float64(result_array[middle_int]+result_array[middle_int-1]) / 2.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tot_len := len(nums1) + len(nums2)
	middle := float64(tot_len) / 2.0
	middle_int := int(middle)
	if float64(middle_int) != middle {
		middle_int = int(math.Floor(middle))
		if len(nums1) > middle_int {
			return float64(nums1[middle_int])
		} else {
			return float64(nums2[middle_int-len(nums1)])
		}
	}
	floor_idx := int(math.Floor(middle))
	ceil_idx := int(math.Ceil(middle))
	floor_val := 0
	ceil_val := 0

	//now we have numbers that need to calculate the mean
	if len(nums1) > floor_idx {
		floor_val = nums1[floor_idx]
	} else {
		floor_val = nums2[floor_idx-len(nums1)]
	}
	if len(nums1) > ceil_idx {
		ceil_val = nums1[ceil_idx]
	} else {
		ceil_val = nums2[ceil_idx-len(nums1)]
	}
	return float64(floor_val+ceil_val) / 2.0

}

func main() {
	fmt.Printf("\n\n %f", findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
	//fmt.Printf("\n\n %f", findMedianSortedArrays2([]int{1, 2}, []int{3, 4, 5}))
	//fmt.Printf("\n\n %f", findMedianSortedArrays2([]int{1, 2, 3}, []int{4, 5}))
}
