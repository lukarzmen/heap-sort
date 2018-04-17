package main

type MaxHeap struct {
	slice    []int
	heapSize int
}

func BuildMaxHeap(slice []int) (maxHeap MaxHeap) {
	maxHeap = MaxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		maxHeap.MaxHeapify(i)
	}
	return
}

func (h MaxHeap) MaxHeapify(index int) {
	leftChildIndex, rightChildIndex := 2*index+1, 2*index+2
	maxValueIndex := index

	if leftChildIndex < h.heapSize && h.slice[leftChildIndex] > h.slice[maxValueIndex] {
		maxValueIndex = leftChildIndex
	}
	if rightChildIndex < h.heapSize && h.slice[rightChildIndex] > h.slice[maxValueIndex] {
		maxValueIndex = rightChildIndex
	}
	if maxValueIndex != index {
		h.slice[index], h.slice[maxValueIndex] = h.slice[maxValueIndex], h.slice[index]
		h.MaxHeapify(maxValueIndex)
	}
}

func HeapSort(slice []int) []int {
	heap := BuildMaxHeap(slice)
	for index := heap.heapSize - 1; index >= 1; index-- {
		heap.slice[0], heap.slice[index] = heap.slice[index], heap.slice[0]
		heap.heapSize--
		heap.MaxHeapify(0)
	}
	return heap.slice
}
