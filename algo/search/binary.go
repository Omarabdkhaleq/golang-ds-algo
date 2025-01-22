package search

import (
	"ds/algo/sort"
)

func IterationBinarySearch(arr []int, target int) (int, bool) {
	arr = sort.QuickSort(arr)
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid, true
		} else if target > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1, false
}

func RecursiveBinarySearch(arr []int, target int, l int, r int) (int, bool) {
	arr = sort.QuickSort(arr)
	if l > r {
		return -1, false
	}
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid, true
		} else if target > arr[mid] {
			return RecursiveBinarySearch(arr, target, mid+1, r)
		} else {
			return RecursiveBinarySearch(arr, target, l, mid-1)
		}
	}
	return -1, false
}
