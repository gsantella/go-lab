package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("writing to a file")

	// data to write in file
	data := []byte("Hello, World!")

	// create or truncate the file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// write data into file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
