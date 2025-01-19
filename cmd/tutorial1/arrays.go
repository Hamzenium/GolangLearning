package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{4, 5, 6}
	intArr[1] = 132
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[2])
}
