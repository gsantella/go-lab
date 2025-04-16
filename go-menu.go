package main

import (
	"bufio"
	"fmt"
	"os"
)

// todo: global data structure ???????????????

// todo: ???????????????????????????????
var scanner = bufio.NewScanner(os.Stdin)

func printMenu() {
	fmt.Println("\n======== MENU =======")
	fmt.Println("  1. Item 1")
	fmt.Println("  2. Item 2")
	fmt.Println("  3. Item 3")
	fmt.Println("  4. Item 4")
	fmt.Print("\nChoice: ")
}

func getUserInput() string {
	scanner.Scan()
	return scanner.Text()
}

func runOption1() {
	fmt.Println("Run Option 1")
}

func runOption2() {
	fmt.Println("Run Option 2")
}

func runOption3() {
	fmt.Println("Run Option 3")
}

func writeDataStructureIntoCSV() {
	fmt.Println("I will do this sometime!")
}

func runCloseProgram() {
	os.Exit(0)
}

func main() {

	//read csv data into the data structure

	for {

		printMenu()
		choice := getUserInput()

		switch choice {
		case "1":
			runOption1()
		case "2":
			runOption1()
		case "3":
			runOption1()
		case "4":
			writeDataStructureIntoCSV()
			runCloseProgram()
		default:
			fmt.Println("Invalid Choice")
		}

	}

}
