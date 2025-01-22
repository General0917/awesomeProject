package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	/*
		defer fmt.Println("world")

		foo()

		fmt.Println("hello")
	*/

	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/

	file, _ := os.Open("./section4/defer/lesson.go")
	defer file.Close()
	data := make(([]byte), 100)
	file.Read(data)
	fmt.Println(string(data))

	/*
		file, err := os.Open("./section4/defer/lesson.go")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		data := make([]byte, 100)
		_, err = file.Read(data)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(string(data))
	*/
}
