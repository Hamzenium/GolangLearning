package main

import "fmt"

func main() {

	var p *int32 = new(int32)
	var i int32

	fmt.Printf("&v", *p)
	fmt.Printf("&v", i)

}
