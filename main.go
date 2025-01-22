package main

import (
	"ds/algo/search"
	"ds/algo/sort"
	"ds/ds/heap"
	"ds/ds/linkedlist"
	"ds/ds/queue"
	"ds/ds/stack"
	"ds/ds/tree"
	"fmt"
)

func main() {

	arr := []int{9, 2, 7, 5, 10, 1, 8, 3, 6, 4}
	// SEARCHING

	fmt.Println("SEARCH")

	//binary
	i, isFound := search.IterationBinarySearch(arr, 1)
	println("Element is found at index", i, isFound)

	i2, isFound2 := search.IterationBinarySearch(arr, 88)
	println("Element is not found", i2, isFound2)

	i3, isFound3 := search.RecursiveBinarySearch(arr, 8, 0, len(arr)-1)
	println("Element is found at index", i3, isFound3)

	i4, isFound4 := search.RecursiveBinarySearch(arr, 88, 0, len(arr)-1)
	println("Element is not found", i4, isFound4)

	// linear
	i5, isFound5 := search.LinearSearch(arr, 8)
	println("Element is found at index", i5, isFound5)

	i6, isFound6 := search.LinearSearch(arr, 88)
	println("Element is not found", i6, isFound6)

	// SORTING

	quickSortedArray := sort.QuickSort(arr)
	fmt.Println("Array sorted using quick sort ", quickSortedArray)

	mergeSortedArray := sort.MergeSort(arr)
	fmt.Println("Array sorted using merge sort ", mergeSortedArray)

	insertionSortedArray := sort.InsertionSort(arr)
	fmt.Println("Array sorted using insertion sort ", insertionSortedArray)

	selectionSortedArray := sort.SelectionSort(arr)
	fmt.Println("Array sorted using selection sort ", selectionSortedArray)

	bubbleSortedArray := sort.BubbleSort(arr)
	fmt.Println("Array sorted using bubble sort ", bubbleSortedArray)

	// DS

	// heap
	heap.HeapDemo()

	// Tree
	tree.BSTDemo()
	tree.BinaryTreeDemo()

	// Linked list
	linkedlist.SinglyLinkedListDemo()
	linkedlist.DoublyLinkedListDemo()

	// Stack
	fmt.Println("Stacks")
	stack.FromScratch()

	// Simple Queue
	fmt.Printf("\n\n")
	fmt.Println("Simple Queue")
	queue.SimpleQueueFromScratch()
	queue.SimpleQueueLinkedList()

	// Circular Queue
	fmt.Printf("\n\n")
	fmt.Println("Circular Queue")
	queue.CircularQueueFromScratch()

	// DQueue
	fmt.Printf("\n\n")
	fmt.Println("DQueue")
	queue.DQueueFromScratch()

}
