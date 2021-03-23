package main

import (
	"fmt"
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
