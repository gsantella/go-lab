package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter your age: ")
	scanner.Scan()
	age := scanner.Text()

	fmt.Print("Enter your city: ")
	scanner.Scan()
	city := scanner.Text()

	fmt.Println("Hello", name, " that is ", age, "years old from", city)

}
