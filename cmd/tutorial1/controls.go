package main

import (
	"errors"
	"fmt"
)

func main() {
	printme()
	var x int = 5
	var y int = 7
	var w, z, err = intDivision(x, y)
	if err != nil {
		fmt.Println(err.Error())
	} else if z == 0 {
		fmt.Println("The result of the integer divison is %v", w)
	} else {
		fmt.Println("The result of the integer divison is %v, and remainder %v", w, z)
	}
}

func printme() {
	fmt.Println("Hello World")
}

func intDivision(numerator int, denonminator int) (int, int, error) {

	var err error
	if denonminator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denonminator
	var remainder int = numerator % denonminator
	return result, remainder, err

}
