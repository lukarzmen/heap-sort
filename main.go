package main

import (
	"fmt"
	"math/rand"
)

func getIntArray(size int) (array []int) {
	array = make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(100)
	}
	return
}

func main() {
	slice := getIntArray(25)
	fmt.Println("Before sorting: ", slice)
	heap := BuildMaxHeap(slice)
	fmt.Println("After build max heap: ", heap.slice)

	slice = HeapSort(slice)
	fmt.Println("After heapSort: ", slice)
}
