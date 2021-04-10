package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var arr = [10]int{1,4,7,8,3,9,2,4,1,8}
	var addedSum int = 18
	var combinations [19]int
	findElementsWithSum(arr,combinations,10,addedSum,0,0,0)
	//fmt.Println(check)//var check2 bool = findElement(arr,9)
	//fmt.Println(check2)
}

func findElementsWithSum(arr[10] int,combinations[19] int,size int, k int, addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num +1
		var p int =0
		for p=0; p < m; p++ {
			time.Sleep(5 * time.Millisecond)
			fmt.Printf("%d,",arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i=l; i< size; i++ {
		//fmt.Println(" m", m)
		combinations[m] = l
		findElementsWithSum(arr,combinations,size,k,addValue+arr[i],l,m+1)
		l = l+1
	}
	return num
}

func bruteForce(arr [10]int, k int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

func linierComplexity()  {
	numsets := []int{10,9,9,7,6,5,4,3,2,1,16}
	bigges := numsets[0]
	for i := 1; i < len(numsets); i++ {
		time.Sleep(time.Second * 1)
		if numsets[i] > bigges{
			log.Println("checking -", i)
			bigges = numsets[i]
			if i == len(numsets)-1 {
				break
			}
		}
		log.Println("checking -", i)
	}
	fmt.Println(bigges)
}
func quadraticComplexity()  {
	for k := 1; k <= 10; k++ {
		for i := 1; i <= 10 ; i++ {
			var x int = 1 * k
			fmt.Println(x)
		}
	}
}
func cubicComplexity()  {
	var arr[10][10][10] int
	for k := 0; k < 10; k++ {
		for l := 0; l < 10; l++ {
			for m := 0; m < 10; m++ {
				arr[k][l][m] = 1
				fmt.Println("Element value ", k, l, m, " is", arr[k][l][m])
			}
		}
	}
}

type Tree struct {
	LeftNode *Tree
	Value int
	RightNode *Tree
}

func (tree *Tree) insert( m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil,m,nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil,m,nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil,m,nil}
	}
}

//print method for printing a Tree
func print(tree *Tree) {
	if tree != nil {
		fmt.Println(" Value",tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		print(tree.RightNode)
	} else {

	}
	fmt.Printf("Nil\n")
}

func logarithmicComplexity()  {
	var tree *Tree = &Tree{nil,1,nil}
	print(tree)

	tree.insert(3)
	print(tree)

	tree.insert(5)
	print(tree)

	tree.LeftNode.insert(7)
	print(tree)
}

