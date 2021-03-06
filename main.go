package main

import (
	"fmt"
	"time"
)

var numbers = []int{9,8,7,6,5,4,3,2,1}
var length = len(numbers)

func main() {
	var ints []int
	for i := 1; i <= 100000; i++ {
		ints = append(ints, i)
	}

	/* Simple Search
		status, index, value := simpleSearch(ints, 500)
		fmt.Println("status : " + status)
		fmt.Println("index : ", strconv.Itoa(index))
		fmt.Println("value : " + strconv.Itoa(value))
	*/

	/* Binary Search
		result, index := binnarySearch(ints, 50000)
		fmt.Println(result, " -> " ,index)
	*/

	/* Bubble Sort
		result := bubbleSort([]int{2,5,4,1,7,12,10})
		for _, val := range result {
			fmt.Println(val)
		}
	*/

	/* Selection Sort
		result := selectionSort(numbers)
		println("\n====================================")
		fmt.Println("hasil akhir : ", result)
		println("====================================")
	*/

	/* recursion
		countDown(10)
		res := recursionEx(10)
		println(res)
	*/

	/*
		result := sum([]int{1,2,3,4,5})
		fmt.Print(result)
		a := fibonacci(5)
		fmt.Println(a)
	*/

	/*
		slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		a := mergeSort(slice)
		fmt.Println(a)
	*/

	/*
		result := quickSort([]int{10,9,8,7,6,5,4,3,2,1})
		fmt.Println(unique(result))
	*/

	/*
		result := insertionSort([]int{10,9,8,7,6,5,4,3,2,-2,-10,-3})
		fmt.Println(result)
	*/

	hashTableSample()
}

func unique(numbers []int) []int{
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range numbers {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func simpleSearch(list []int, item int)(status string, index int, value int) {
	for i, val := range list {
		if item == val {
			status = "found"
			index = i
			value = val
			return
		}
		fmt.Println("scanning...")
	}

	status = "not found"
	index = -1
	value = -1
	return
}

func binnarySearch(list []int, item int) (int, int) {
	low := 0
	high := len(list) - 1
	for {
		if low <= high {
			mid := (low + high) / 2
			guess := list[mid]
			if guess == item {
				return item, mid
			}
			if guess > item {
				high = mid - 1
			}else {
				low = mid + 1
			}
		}
	}
	return -1, -1
}

func bubbleSort(list []int) (sorted []int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if (list[j] > list[j+1]) {
				temp := list[j]
				list[j] = list[j+1];
				list[j+1] = temp;
			}
		}
	}
	sorted = list
	return
}

func findSmallest(numbers []int) int {
	smallest := numbers[0]
	smallest_index := 0	

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
			smallest_index = i
		}	
	}
	return smallest_index
}

func selectionSort(numbers[] int) []int{
	var newArr []int 
	for i := 0; i < length; i++ {
		smallest := findSmallest(numbers)
		newArr = append(newArr, numbers[smallest])
		
		println("----------------------------------------------------------------------")
		fmt.Println("smallest index ke : ", smallest)
		fmt.Println("value yg diambil  : ", numbers[smallest])
		fmt.Println("slice baru        : ", newArr)
	
		numbers = remove(numbers, smallest)
		fmt.Println("numbers setelah di remove index ke  : ",smallest, " : ", numbers)
	}
	return newArr
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func countDown(x int) int {
	if x < 1 {
		println("time out")
		return -1
	}
	time.Sleep(100 * time.Millisecond)
	println(x)
	return countDown(x-1)
}

func recursionEx(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursionEx(n-1)
}

func sum(numbers []int) int {
	temp := 0
	for _, v := range numbers {
		temp += v
	}
	return temp
}

func fibonacci(n int)  int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2

	// [10 9 8 7 6]
	// [5 4 3 2 1]
	return Merge(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
}

func Merge(left, right []int) []int {

	// size = 10
	// i = 0
	// j = 0
	size, i, j := len(left)+len(right), 0, 0

	slice := make([]int, size, size)

	// [10 9 8 7 6]
	// [5 4 3 2 1]
	for k := 0; k < size; k++ {
		if  i > len(left) - 1 && j <= len(right) - 1{
			slice[k] = right[j]
			j++
		}else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		}else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		}else {
			slice[k] = right[j]
			j++
		}
	}

	return slice
}

func quickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	var pivot []int 
	pivot = append(pivot, numbers[0])
	var less []int
	var greater []int

	for _, v := range numbers {
		if v < pivot[0] {
			less = append(less, v)
		}else {
			greater = append(greater, v)
		}
	}

	a := append(quickSort(less), pivot...)
	return append(a,quickSort(greater)...)
}

func insertionSort(numbers []int) (sorted []int) {
	temp := numbers[:]
	for j := 1; j < len(temp); j++ {
		key, i := temp[j], j-1
		for i > -1 && temp[i] > key {
			temp[i+1], i = temp[i], i-1
		}
		temp[i+1] = key
	}
	sorted = temp
	return
}

func hashTableSample() {
	var _ map[string]interface{}
}