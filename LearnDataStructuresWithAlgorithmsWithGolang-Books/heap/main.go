package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iHeap IntegerHeap) Len() int {
	return len(iHeap)
}

func (iHeap IntegerHeap) Less(i,j int) bool {
	return iHeap[i] < iHeap[j]
}

func (iHeap IntegerHeap) Swap(i,j int)  {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

func (iHeap *IntegerHeap) Push(heapintf interface{})  {
	*iHeap = append(*iHeap, heapintf.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1,4,5,0}
	intHeap.Push(10)
	intHeap.Push(11)
	intHeap.Push(12)

	heap.Init(intHeap)
	fmt.Printf("minimum %v", (*intHeap)[0])
	fmt.Println("\n\n")

	intHeap.Pop()
	heap.Push(intHeap, 2)
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}