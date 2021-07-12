package main

import "fmt"

func Comment() {
	// this is a single line comment
	/*
	this is a multi line comment
	*/
}

func DataTypes() {
	// Declaring and Initializing Constant
	const con1 = 1
	const con2 = 15

	// Declaring and Initializing Variable
	var var1 = 5
	var var2 = 12

	// Changing Variable value
	var1 = 25
	var2 = 35

	// Print Value
	fmt.Printf("con1 is %d \n", con1)
	fmt.Printf("con2 is %d \n", con2)
	fmt.Printf("var1 is %d \n", var1)
	fmt.Printf("con2 is %d \n", var2)

}

func Constants()  {

	// Declaring and Initializing Constant
	const con1 = 5
	const con2 = 20

	fmt.Println("First Constant :", con1)
	fmt.Println("Second Constant :", con2)

}

func testVariables()  {
	// Declaring and Initializing Variable
	var var1 = 5
	var var2 = 19

	// Changing Variable Value
	var1 = 20
	var2 = 30

	// Print Values
	fmt.Println("First Variable :", var1)
	fmt.Println("Second Variable :", var2)
}

func main() {
	//DataTypes()
	//Constants()
	testVariables()
}
