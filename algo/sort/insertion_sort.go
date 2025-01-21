package sort

func InsertionSort(arr []int) []int {
	for i, v := range arr {
		j := i - 1
		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
	}
	return arr
}
