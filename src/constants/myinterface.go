package constants

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodwithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters call")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithoutParameters call", f)
}

func (m MyType) MethodwithReturnValue() string {
	return "Hi from MethodWithoutParameters"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface call")
}
