/*
Two Sum

Given an array of integers nums and an integer target, return indices
 of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

// A Brute Force Solution — O(N²) Time, O(1) Space
package main

import "sort"

func twoSum_brute_force(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		var start = array[i]
		for j := i + 1; j < len(array); j++ {
			var end = array[j]
			if (start + end) == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// Using a Hash Map — O(N) Time, O(N) Space

func twoSum_hash(array []int, target int) []int {
	seenNums := map[int]int{}
	for i, num := range array {
		potentialMatch := target - num
		if j, found := seenNums[potentialMatch]; found {
			return []int{i, j}
		}
		seenNums[num] = i
	}
	return []int{}
}

// Sorting & Crunching— O(NlogN) Time, O(1) Space

func twoSum_sort(array []int, target int) []int {
	sort.Ints(array)
	start, end := 0, len(array)-1
	for start < end {
		currentSum := array[start] + array[end]
		if currentSum == target {
			return []int{start, end}
		} else if currentSum < target {
			start += 1
		} else {
			end -= 1
		}
	}
	return []int{}
}
