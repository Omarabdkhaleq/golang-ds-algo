package main

import (
	"ds/algo/sort"
	"fmt"
)

func main() {

	arr := []int{9, 2, 7, 5, 10, 1, 8, 3, 6, 8}

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

	//heap.HeapDemo()

	//tree.BSTDemo()
	//tree.BinaryTreeDemo()

	//linkedlist.SinglyLinkedListDemo()
	//linkedlist.DoublyLinkedListDemo()

	//// Stack
	//fmt.Println("Stacks")
	//stack.FromScratch()

	//// Simple Queue
	//fmt.Printf("\n\n")
	//fmt.Println("Simple Queue")
	//queue.SimpleQueueFromScratch()
	//queue.SimpleQueueLinkedList()

	//// Circular Queue
	//fmt.Printf("\n\n")
	//fmt.Println("Circular Queue")
	//queue.CircularQueueFromScratch()

	//// DQueue
	//fmt.Printf("\n\n")
	//fmt.Println("DQueue")
	//queue.DQueueFromScratch()

}
