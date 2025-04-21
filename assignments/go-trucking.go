package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var zips []int

	for {

		var scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		zipAsString := scanner.Text()

		zipAsInt, _ := strconv.Atoi(zipAsString)

		// exit on sentinel value < 0
		if zipAsInt < 0 {
			break
		}

		zips = append(zips, zipAsInt)

	}

	sort.Ints(zips)

	for _, val := range zips {
		fmt.Printf("%05d\n", val)
	}

}
