package main

import "fmt"

//function to solve the first project euler problem
func multiple(a int, b int) int {
	list := []int{}
	for i := 0; i < 1000; i++ {
		if i%a == 0 {
			list = append(list, i)
		}
		if i%b == 0 {
			list = append(list, i)
		}
	}
	removeDuplicateInt(list)
	sum := 0
	for _, num := range list {
		sum += num
	}
	return sum
}

//function from https://stackoverflow.com/questions/66643946/
//how-to-remove-duplicates-strings-or-int-from-slice-in-go
//to remove duplicate ints on slice
func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

//main to call and print the functions above
func main() {
	fmt.Println("The first Project Euler problem answer:")
	fmt.Println(multiple(3, 5))
}
