// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-13
// Fileoverview: This program tells you if you should eat based on hunger level.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var hungerAsString string = ""
	var hungerAsNumber int = 0

	var reader = bufio.NewReader(os.Stdin)

	// input hunger level
	fmt.Print("Hello, how hungry are you on a scale of 1 to 10, where 1 is not hungry and 10 means you must eat? ")
	hungerAsString, _ = reader.ReadString('\n')
	hungerAsString = strings.TrimSpace(hungerAsString)
	hungerAsNumber, _ = strconv.Atoi(hungerAsString)

	// check hunger level
	if hungerAsNumber > 5 {
		fmt.Println("Please eat!")
	} else {
		fmt.Println("You are not really that hungry.")
	}

	fmt.Println("\nDone.")
}
