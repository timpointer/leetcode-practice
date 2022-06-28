package main

func searchInsert(nums []int, target int) int {

	return searchInsertI(nums, 0, len(nums)-1, target)
}

func searchInsertI(nums []int, start, end int, target int) int {
	if end < start {
		return end + 1
	}

	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return searchInsertI(nums, start, mid-1, target)
	} else {
		return searchInsertI(nums, mid+1, end, target)
	}

	panic("no run")
}
