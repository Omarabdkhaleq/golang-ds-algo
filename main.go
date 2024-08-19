package main

import (
	"ds/ds/queue"
	"ds/ds/stack"
	"fmt"
)

func main() {

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
