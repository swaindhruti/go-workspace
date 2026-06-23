package main

import (
	"fmt"
)

/*
let there be 4 stairs and you are standing  at the bottom and there is some package at the 4th stair you want and get back to the bottom stair you will climb and return so you will have total 8 stairs to climb. but if we have a person at the 4th stair and you are at the bottom stair you will climb up 2 stair the top man will climb down 2 stair and you will meet in between and the thing will be handover and you both will return to respective position here the net work is 8 stairs but the time taken is halfed as befor you were climing 8 stairs nor we are climbing 4 stairs.

Where to can use 2 pointers pattern?
-> question will be of array olr linked list not in other cases in majority of the cases.
-> if sorted data or sorted data we will be easy e can apply this technique.
-> if we have to merge/ remove duplicates/ rearrange/ detect cycle we can use this technique.
-> if ewe have to find triplets, quadruple, pair we can use this technique.
-> Do not use extra space(memory)

2 Pointer and Sliding Window are similar but not the same pattern.
*/

// Sum of array - using 2 pointer pattern
func SumOfArray(arr []int) int {
	left, right := 0, len(arr)-1
	sum := 0

	for left <= right { // <= because we have to consider the case when there is odd number of elements in the array and left and right will be pointing to the same element
		if left == right {
			sum += arr[left]
		} else {
			sum += arr[left] + arr[right]
		}
		left++
		right--
	}

	return sum
}

// 2 Sum of Array - Find two no.s in the array that add up to a target sum - using 2 pointer pattern

// we can use this technique where the elements of the arr is asked or index is asked on an sorted array. If index is asked we cannot do the sorting as it will change the required result.

func TwoSum(arr []int, target int) [2]int {
	left, right := 0, len(arr)-1

	for left < right {
		currentSum := arr[left] + arr[right]

		if currentSum == target {
			return [2]int{arr[left], arr[right]}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return [2]int{-1, -1} // Return -1, -1 if no such pair is found
}

// Remove duplicates from sorted array
func RemoveDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	left := 0
	for right := range len(arr) {
		if arr[right] != arr[left] {
			left++
			arr[left] = arr[right]
		}
	}

	return arr[:left+1]
}

// Leetcode : Remove duplicates from sorted array 2
func RemoveDuplicate2(arr []int) []int {
	if len(arr) < 3 {
		return arr
	}

	left := 2
	for right := 2; right < len(arr); right++ {
		if arr[right] != arr[left-2] {
			arr[left] = arr[right]
			left++
		}
	}

	return arr[:left]
}

func SortedSquares(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Sort by absolute value
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			// Get absolute values
			absI := arr[i]
			if absI < 0 {
				absI = -absI
			}
			absJ := arr[j]
			if absJ < 0 {
				absJ = -absJ
			}

			// Swap if left's absolute value is greater
			if absI > absJ {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sumOfArrayResult := SumOfArray(arr)
	fmt.Println(sumOfArrayResult)

	TwoSumResult := TwoSum(arr, 7)
	fmt.Println(TwoSumResult[0], TwoSumResult[1])

	arr2 := []int{1, 1, 1, 1, 3, 3, 3, 3, 6, 777, 777, 4545}
	RemoveDuplicatesResult := RemoveDuplicates(arr2)
	fmt.Println(RemoveDuplicatesResult)

	arr3 := []int{1, 1, 1, 1, 3, 3, 3, 3, 6, 777, 777, 4545}
	RemoveDuplicate2Result := RemoveDuplicate2(arr3)
	fmt.Println(RemoveDuplicate2Result)

	arr4 := []int{-4, -3, -2, 0, 3, 4}
	SquareOfSortedArray := SortedSquares(arr4)
	fmt.Println(SquareOfSortedArray)
}
