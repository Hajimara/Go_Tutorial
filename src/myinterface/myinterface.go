package main

import (
	"constants"
	"fmt"
)

func main() {
	var value constants.MyInterface
	value = constants.MyType(5)
	value.MethodWithParameter(5)
	value.MethodWithoutParameters()
	fmt.Println(value.MethodwithReturnValue())
}
