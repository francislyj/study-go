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

func hello(data string, channelCall chan string) {
	fmt.Println("channel called")
	data = data + "hello"
	fmt.Println("data changed")

	fmt.Println("channel sleeping for 2 seconds")
	time.Sleep(2 * time.Second)
	fmt.Println("channel awake")

	channelCall <- data
	fmt.Println("channel message send :", data)
}

func main() {
	//var j  = 1
	//j ++
	//for i := 0; i < 3; i ++ {
	//	go print(i)
	//}
	//
	//var input string
	//fmt.Scanln(&input)

	// channel
	data := "one percent"
	channelCall := make(chan string)

	go hello(data, channelCall)

	receivedData := <- channelCall

	fmt.Println("received message data :", receivedData)

}
