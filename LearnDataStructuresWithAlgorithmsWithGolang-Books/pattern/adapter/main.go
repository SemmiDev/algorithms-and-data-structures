package main

import "fmt"

type (
	IProcess interface { process() }

	Adapter struct { adaptee Adaptee }
	Adaptee struct {
	finally Finally
	PID     int
}
	Finally struct { exit int }
)

func (adapter Adapter) process() {
	fmt.Println("adapter process")
	adapter.adaptee.convert()
}

func (adaptee Adaptee) convert() {
	fmt.Println("adaptee convert method")
	adaptee.PID = 100
	adaptee.finally.killProcess(adaptee.PID)
}

func (finally Finally) killProcess(PID int) {
	finally.exit = PID
	fmt.Println("kill process ", finally.exit)
}

func main() {
	//var processor IProcess = Adapter{}
	//processor.process()

	//finally := Finally{}
	//adaptee := Adaptee{finally: finally, PID: 100}
	//adapter := Adapter{adaptee: adaptee}
	//adapter.process()
}
