package main

import "fmt"

func multiple(a int) []int {
	list := []int{}
	for i := 0; i < 1000; i++ {
		if i%a == 0 {
			list = append(list, i)
		}
	}
	return list
}

func main() {
	fmt.Println(multiple(3))
}
