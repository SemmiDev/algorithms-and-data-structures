package main

import (
	"fmt"
)

func main() {
	var ints []int
	for i := 1; i <= 100000; i++ {
		ints = append(ints, i)
	}

	/* SimpleSearch
		status, index, value := simpleSearch(ints, 500)
		fmt.Println("status : " + status)
		fmt.Println("index : ", strconv.Itoa(index))
		fmt.Println("value : " + strconv.Itoa(value))
	*/

	/* BinarySearch
		result, index := binnarySearch(ints, 50000)
		fmt.Println(result, " -> " ,index)
	*/

	result := bubbleSort([]int{2,5,4,1,7,12,10})
	for _, val := range result {
		fmt.Println(val)
	}
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