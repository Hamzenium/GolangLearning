package main

import "fmt"

func main() {
	fmt.Println("Hello Momina Mustehsan")
	var intNum int
	var floatNum32 float32 = 10.2
	var intNum32 int32 = 2
	fmt.Println(intNum)

	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var mystring string = "hello" + " " + "world"
	fmt.Println(mystring)
}
