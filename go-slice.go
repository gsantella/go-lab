package main

import "fmt"

func main() {

	// arrays
	//var nums [10]int
	//var num2 = [10]int{6, 5, 4, 6, 5, 4, 3, 2, 1, 0}
	//var num3 = [...]int{1, 2, 3, 4}

	// slices
	//var slice1 []int
	var slice2 = []int{1, 2, 3, 4}

	slice2 = append(slice2, 5, 4, 3, 2, 1)
	slice2 = slice2[:2]

	for _, value := range slice2 {
		fmt.Println(value)
	}

	fmt.Println("csv")
}
