package main

import "fmt"

func main() {
	//1st usul of declaring a variable
	//var age int = 24 //variable declaration with initial value
	//2nd usul of declaring a variable
	//var age = 24
	//fmt.Println("My initial age is", age)
	
	//Multiple variable declaration
	var day, month, year int = 29, 04, 2026 //declared multiple variable
	fmt.Println("day is", day, "month is", month, "year is", year)
	//I wanna see formatted version of above example, let's go
	fmt.Printf("Date: %02d%02d%d\n", day, month, year)
	fmt.Printf("Date: %02d.%02d.%d\n", day, month, year)
}
