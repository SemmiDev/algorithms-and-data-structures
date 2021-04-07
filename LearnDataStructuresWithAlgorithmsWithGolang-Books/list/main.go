package list

import (
	"container/list"
	"fmt"
)

type Student struct {
	ID 		string
	Name 	string
}

func thisIsList()  {
	var students = []Student{
		{"1","sam"},
		{"2","dev"},
		{"3","sammidev"},
	}

	var studentLIST list.List

	studentLIST.PushBack(students[0])
	studentLIST.PushBack(students[1])
	studentLIST.PushBack(students[2])
	studentLIST.PushFront(Student{ID: "0",Name: "sammidev"})

	for element := studentLIST.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(Student))
	}

	enter()


	frontValue := studentLIST.Front().Value.(Student)
	backValue := studentLIST.Back().Value.(Student)

	fmt.Println(frontValue)
	fmt.Println(backValue)
}

func enter() {
	fmt.Println("\n\n\n")
}

