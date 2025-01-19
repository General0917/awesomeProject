package main

import "fmt"

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f         = true, false
)

func foo() {
	xi := 1
	//xf64 := 1.2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

func main() {
	/*
		var i int = 1
		var f64 float64 = 1.2
		var s string = "test"
		//var t bool = true
		//var f bool = false
		var t, f = true, false
	*/
	fmt.Println(i, f64, s, t, f)
	foo()
}
