package main

import (
	"container/list"
	"fmt"
)

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

func testInOut() {
	// Display a message
	fmt.Println("One Percent")
	fmt.Println("Enter your Name :")

	var name string

	// Asks for Input
	// Stored the Input in the Variable "name"
	fmt.Scanln(& name)

	fmt.Println("Welcome to the One Percent Club :", name)
}

func main() {
	//DataTypes()
	//Constants()
	//testVariables()
	//testInOut()

	var str = "aaa" + "bbb"
	result := fmt.Sprintf("hello %s", str)
	fmt.Println(result)

	var a complex64 = 28
	fmt.Println(a)

	var b int8 = 8
	var c float32 = float32(b)
	fmt.Println(c)

	var arr1 []int = []int{1, 2, 3}
	fmt.Println(arr1)



	var p1 *int
	p1 = new(int)
	//*p1 = 10
	fmt.Println(*p1)

	var myList = []string{"🍌", "🍉"}
	fmt.Println(myList)

	h := list.New()
	h.PushFront("aa")

	r := &Rect{width: 20, height: 30}
	hello2(r)
}

func hello2(ir IR) {
	fmt.Println(ir.getWidth())
}

type IR interface {
	getWidth() int
}

type Rect struct {
	width int
	height int
}

func (r Rect) getWidth() int {
	return r.width
}

func Add(num1, num2 int) int {
	return num1 + num2
}