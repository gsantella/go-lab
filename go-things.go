package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type die struct {
	value uint8
}

func (d die) String() string {
	return fmt.Sprintln("A great die with value", d.value)
}

func (d *die) roll() {
	d.value = uint8(rand.IntN(6) + 1)
}

func (d *die) setValue(value uint8) error {

	if value < 1 || value > 6 {
		return errors.New("Invalid die value")
	}

	d.value = value
	return nil

}

func main() {
	d1 := die{}

	if err := d1.setValue(2); err != nil {
		fmt.Println("Bad Value")
	} else {
		fmt.Print(d1)
	}

	d1.roll()
	d1.roll()
	d1.roll()

	fmt.Print(d1)
}
