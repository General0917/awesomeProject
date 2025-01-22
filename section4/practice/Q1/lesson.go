package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int

	for i, num := range l {
		if i == 0 {
			min = num
			continue
		}

		if min > num {
			min = num
		}
	}

	fmt.Println(min)

	// 以下でも同じ結果が得られる
	/*
		m := l[0]
		for _, v := range l {
			if m > v {
				m = v
			}
		}
		fmt.Println(m)
	*/
}
