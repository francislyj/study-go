package main

import "fmt"

func main() {

	// int
	var x1 int = 100
	var x2 = 200
	x3 := 300

	// float
	var x4 = 2.0
	x5 := 54.32

	// bool
	var x6 = true
	x7 := false

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)
	fmt.Println(x6)
	fmt.Println(x7)


	// array
	var students = [5]string{"d", "sss", "sss"}

	for k, v := range students {
		fmt.Println(k, v)
	}

	// string
	var str = "abcde"
	fmt.Println(str)
	for k, v := range str {
		fmt.Println(k, v)
	}

	// pointer
	var a1 = "aa"
	var pa1 *string = &a1
	fmt.Println(a1)
	fmt.Println(pa1)
	*pa1 = "bb"
	fmt.Println(a1)

	// Structure
	type Employee struct {
		ID string
		Name string
	}

	em1 := Employee{
		ID: "a1",
		Name: "N1",
	}

	em2 := Employee{
		ID: "A2",
		Name: "N2",
	}

	var ems = []*Employee{&em1, &em2}
	for i, em := range ems {
		fmt.Println(i, em.ID, em.Name)
	}




}
