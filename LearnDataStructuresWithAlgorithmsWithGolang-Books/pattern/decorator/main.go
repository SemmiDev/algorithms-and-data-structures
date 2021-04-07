package main

import (
	"fmt"
)

type IProcess interface {
	process()
}

//ProcessClass struct
type ProcessClass struct{}

//ProcessClass method process
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

//ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

//ProcessDecorator class method process
func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

// // main method
//func main() {
//	var process = &ProcessClass{}
//	var decorator = &ProcessDecorator{}
//
//	decorator.process()
//
//	decorator.processInstance = process
//	decorator.process()
//}

type SamProcess interface {
	toHome()
	signUp() bool
}

type Home struct {}

func (home *Home) toHome()  {
	fmt.Println("HOME PAGE")
}

type HomeDecorator struct {
	homeInstance *Home
}

func (decorator *HomeDecorator) toHome()  {
	if decorator.homeInstance == nil {
		fmt.Println("1 process")
	}else {
		fmt.Print("1 process and ")
		decorator.homeInstance.toHome()
	}
}

func main() {
	home := Home{}
	decorator := HomeDecorator{}
	decorator.toHome()

	decorator = HomeDecorator{homeInstance: &home}
	decorator.toHome()
}