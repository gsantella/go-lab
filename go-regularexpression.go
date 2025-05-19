package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// generate a uuid, verify with regexp, strings

func main() {
	fmt.Println("Regular Expressions!")
	fmt.Println("Give me some text: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)

	// function chaining (fluent pattern) not a default in go
	//text.ToUpper().TrimSpace()

	fmt.Println(text)

	// Regular Expression Thing
	re := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)
	matches := re.FindAllString(text, -1)
	fmt.Println("Found Emails: ", matches)

	/*take input
	  show count before trim
	  trim both sides and also make sure all multiple spaces are removed
	  show count after trim
	  find your name in string
	  		if there show what index it starts with
	  			then replace your name with FOUND HERE printing result
	  		else find that all the letters exist in the string of you name
	*/

	// 2 funcs - generate a uuid, verify with regexp
	// test it out with some sample code - your uuid generator
	// should verify every time
}
