package main

import (
    "fmt"
    "os"
	"strconv"
	"strings"
)

var digits = []string{"Zero", "One", "Two", "Three", "Four", "Five",
					  "Six", "Seven", "Eight", "Nine",}

func main() {
	// Get all the input arguments
	args := os.Args[1:]
	
	var output []string
	// Loop through each argument
	for i := 0; i < len(args); i+=1 {
		arg := args[i]
		// If the argument is not a number => Throw error and continue
		_, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println(arg + " is not a valid integer.")
			continue
		}
		str := ""
		// If argument is a negative integer, add "Minus" in front
		start := 0
		if string(arg[0]) == "-" {
			start = 1
			str += "Minus"
		}
		// Loop through each digit
		for j := start; j < len(arg); j+=1 {
			// Convert to string
			number, _ := strconv.Atoi(string(arg[j]))
			// Add each phonetic of a digit
			str += digits[number]
		}
		// Print out the number
		output = append(output, str)
	}
	if output != nil {
		fmt.Println(strings.Join(output, ","))
	}
}