package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var a int = 5
	fmt.Println(a)
	//file, err := os.Create("hello.txt")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("file created success", file)

	file, err := os.Open("hello.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("hello world")

	defer file.Close()
	fmt.Println("file written success")
}
