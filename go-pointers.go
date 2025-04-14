package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	cup := 1
	fmt.Println("initial:", cup)

	zeroval(cup)
	fmt.Println("zeroval:", cup)

	zeroptr(&cup)
	fmt.Println("zeroptr:", cup)

	fmt.Println("pointer:", &cup)
}
