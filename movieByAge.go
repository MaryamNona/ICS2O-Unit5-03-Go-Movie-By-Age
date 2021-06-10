// Created by: Maryam Nona
// Createdon: May 2021
//
// This program does basic math

package main

import (
	"fmt"
)

func main() {
	var userAge int

	// input
	fmt.Println("This program tells you what type of movie you can watch.")
	fmt.Println()
	fmt.Println("Enter you age: ")
	fmt.Scanln(&userAge)

	// output
	if userAge > 17 {
		fmt.Println("You can watch R movies alone.")
	} else if userAge > 13 {
		fmt.Println("You can watch PG-13 Movies alone.")
	} else if userAge > 5 {
		fmt.Println("You can watch G or PG movies alone.")
	} else {
		fmt.Println("You're too young.")
	}
}
