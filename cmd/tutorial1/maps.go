package main

import "fmt"

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	for name := range myMap2 {
		fmt.Printf("name: %v\n", name)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
