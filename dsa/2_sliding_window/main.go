package main

import "fmt"

/*
The window can expland and contract based on the condition, and we can use two pointers to keep track of the window's start and end. The window will always move in one direction, and we can use the two pointers to keep track of the current window's size and position. The window can be of fixed size or variable size based on the problem requirements. Used for array and strings.

Applicable when subarray or substring is asked in the question. Are we finding something maximum or minimum, longest, shortest, sum/count/average, atmost k, atleast k or exactly k.

2 types of sliding window:
1. Fixed size sliding window: The window size is fixed and does not change during the process. We can use this technique when we know the size of the window beforehand.

2. Variable size sliding window: The window size can change during the process based on the condition. We can use this technique when we do not know the size of the window beforehand.
*/

// Maximum Sum Subarray of Size K - Fixed size sliding window
func MaxSumSubarray(arr []int, k int) {
	if len(arr) < k {
		fmt.Println("Array length is less than k")
		return
	}

	maxSum := 0
	windowSum := 0

	for i := range k {
		windowSum += arr[i]
	}
	maxSum = windowSum

	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k] // slide the window forward by adding the next element and removing the first element of the previous window
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	fmt.Println("Maximum sum of subarray of size", k, "is", maxSum)
}

// Minimum Size Subarray Sum - Variable size sliding window
func MinSizeSubarraySum(arr []int, target int) {
	minLength := len(arr) + 1
	windowSum := 0
	left := 0

	for right := range arr {
		windowSum += arr[right]

		for windowSum >= target {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}
			windowSum -= arr[left]
			left++
		}
	}

	if minLength == len(arr)+1 {
		fmt.Println("No subarray found with sum greater than or equal to", target)
	} else {
		fmt.Println("Minimum length of subarray with sum greater than or equal to", target, "is", minLength)
	}
}

func main() {
	MaxSumSubArrayInput := []int{2, 1, 5, 1, 3, 2}
	MaxSumSubarray(MaxSumSubArrayInput, 3)

	MinSizeSubarraySumInput := []int{2, 3, 1, 2, 4, 3}
	MinSizeSubarraySum(MinSizeSubarraySumInput, 7)
}
