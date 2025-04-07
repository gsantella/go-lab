package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("writing to a file")

	// data to write in file
	data := []byte("Hello, World!")

	// create or truncate the file
	file, err := os.Create("C:\\output.txt")
	checkError(err)

	// write data into file
	_, err = file.Write(data)
	checkError(err)
}
