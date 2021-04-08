package main

import "fmt"

type IRealObject interface {
	performAction()
}

type RealObject struct{}

func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

//VirtualProxy struct
type VirtualProxy struct {
	realObject *RealObject
}
//VirtualProxy class method performAction
func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()")
	virtualProxy.realObject.performAction()
}
// main method
func main() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}