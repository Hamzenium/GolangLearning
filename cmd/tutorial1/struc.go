package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func main() {

	var myEngine gasEngine = gasEngine{mpg: 12, gallons: 34}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total Miles: %v", myEngine.milesLeft())
}
