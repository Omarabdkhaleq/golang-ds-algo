package sort

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minValue := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minValue] > arr[j] {
				minValue = j
			}
		}
		arr[i], arr[minValue] = arr[minValue], arr[i]
	}
	return arr
}
