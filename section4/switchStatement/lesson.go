package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "linux"
}

func main() {
	//os := getOsName()

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac OS")
	case "windows":
		fmt.Println("Windows OS")
	default:
		fmt.Println("Unknown OS", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	}
}
