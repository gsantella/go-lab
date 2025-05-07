package main

import (
	"fmt"
)

type Ring struct {
	values []int
	pos    int
}

func (r *Ring) NextValue() int {
	value := r.values[r.pos]
	r.pos = (r.pos + 1) % len(r.values)
	return value
}

func NewRing(values []int) *Ring {
	return &Ring{values: values}
}

func main() {

	// Make a Ring
	numbers := []int{7, 2, 4}
	ring := NewRing(numbers)

	for {
		fmt.Println(ring.NextValue())
	}

	// Ring Algorithm Basics
	values := []int{1, 2, 3, 4, 5}
	pos := 0

	for {
		if pos >= 5 {
			pos = 0
		}
		fmt.Println(values[pos])
		pos += 1
	}

	fmt.Println("Round Robin")
}
