package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print(n int) {
	for i := 0; i < 3; i ++ {
		fmt.Println(n, ":", i)
		randomNum := time.Duration(rand.Intn(250))
		fmt.Println("randomNum :", randomNum)
		time.Sleep(time.Millisecond * randomNum)
	}
}

func main() {
	var j  = 1
	j ++
	for i := 0; i < 3; i ++ {
		go print(i)
	}

	var input string
	fmt.Scanln(&input)

}
