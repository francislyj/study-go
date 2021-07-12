package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {

	var x = 20
	var y = 20

	var sum = x + y
	var difference = x - y
	var product = x * y
	var quotient = x / y
	var remainder = x % y

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)
	fmt.Println(remainder)
	x -= y
	//xixi: fmt.Println(x & y)

	if (x < y) {
		fmt.Println("hehe")
	}

	// goto
	back: for i := 0; i < 5; i ++ {
		fmt.Println("hello", i)
		if i == 4 {
			goto huhu
		}

		if i == 6{
			goto back
		}
	}

	fmt.Println("hello world")

	huhu: fmt.Println("soga")


	// fallthrough
	var abc = 3

	switch abc {
	case 1:
		fmt.Println("hello 1")
	case 2:
		fmt.Println("hello 2")
	case 3:
		fmt.Println("hello 3")
		fallthrough
	case 4:
		fmt.Println("hello 4")
	case 5:
		fmt.Println("hello 5")
	default:
		fmt.Println("hello default")
	}

	// array
	arr1 := []string{"a", "b"}
	arr2 := [][]string{{"cd", "e"}, {"he", "ha"}}
	fmt.Println(arr1)
	for _, item := range arr2 {
		fmt.Println(item)
	}

	// pointer
	var mPointer *int
	var num = 12
	mPointer = &num

	fmt.Println("value of num :", num)
	fmt.Println("value of num :", *mPointer)
	fmt.Println("address of num :", &num)
	fmt.Println("address of num :", mPointer)
	fmt.Println("address of pointer mPointer is:", &mPointer)


	var num1, num2 int
	var ptr1, prt2 *int

	num1 = 10
	num2 = 20

	fmt.Println("before swapping:")
	fmt.Println("num1 is :", num1)
	fmt.Println("num2 is :", num2)
	fmt.Println()

	ptr1 = &num1
	prt2 = &num2

	//temp = *ptr1
	//*ptr1 = *prt2
	//*prt2 = temp
	*ptr1, *prt2 = *prt2, *ptr1

	fmt.Printf("after swapping:")
	fmt.Println("num1 is :", num1)
	fmt.Println("num2 is :", num2)

	// math
	var num3 = -100.02
	var num4 = -225

	var num5 = math.Abs(num3)
	var num6 = math.Abs(float64(num4))
	var num7 = math.Sqrt(9)
	var num8 = math.Pow(3, 4)
	var num9 = math.Round(3.49)
	var num10 = math.Sin(1)

	fmt.Println("absolute value :", num5)
	fmt.Println("absolute value :", num6)
	fmt.Println("absolute value :", num7)
	fmt.Println("absolute value :", num8)
	fmt.Println("absolute value :", num9)
	fmt.Println("absolute value :", num10)

	// string
	str := "abcdaaab"
	fmt.Println(len(str))
	fmt.Println(strings.Count(str, "ab"))
	fmt.Println(strings.Contains(str, "ab"))

	//fmt.Println("please enter num1 :")
	//fmt.Scanln(&num1)
	//fmt.Println("please enter num2 :")
	//fmt.Scanln(&num2)
	//divide(num1, num2)
	fmt.Println("hehe")

	// defer
	defer fun1()
	fun2()


	// Map
	Capitals := map[string]string{"India": "new","hehe": "alal"}

	for key, value := range Capitals{
		fmt.Println(key, value)
	}

	// time
	currentTime := time.Now()
	secs := currentTime.Unix()
	nacos := currentTime.UnixNano()

	millis := nacos/ 1000000
	mins := secs / 60
	hours := mins / 60
	days := hours / 24
	years := days / 365
	months := years * 12

	fmt.Println("Time since epoch :")
	fmt.Println("Nanoseconds :", nacos)
	fmt.Println("milliseconds :", millis)
	fmt.Println("seconds :", secs)
	fmt.Println("minutes :", mins)
	fmt.Println("hours :", hours)
	fmt.Println("days :", days)
	fmt.Println("months :", months)
	fmt.Println("years :", years)

	fmt.Println("test :", currentTime.Month())

	//var i = 0
	//for _ = range time.Tick(1 * time.Second) {
	//	i ++
	//	haveFun(i)
	//}


}

func haveFun(num int) {
	fmt.Println("execution number:", num)
}

func fun1() {
	fmt.Println("fun1 executed")
}

func fun2() {
	fmt.Println("fun2 executed")
}

func divide(x, y int) int {
	defer func() {
		fmt.Println("Errors :", recover())
	}()
	var result = x / y
	fmt.Println("result is :", result)
	return result
}