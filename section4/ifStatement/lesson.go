package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "OK"
	} else {
		return "NO"
	}
}

func main() {
	result := by2(10)

	if result == "OK" {
		fmt.Println("GREAT")
	}
	fmt.Println(result)

	if result2 := by2(10); result2 == "OK" {
		fmt.Println("GREAT 2")
	}
	/*
		num := 6
		if num % 2 == 0 {
			fmt.Println("by 2")
		} else if num % 3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	x, y := 11, 12
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}
