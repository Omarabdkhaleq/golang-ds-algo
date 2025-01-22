package search

func LinearSearch(arr []int, target int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i, true
		}
	}
	return -1, false
}
