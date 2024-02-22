package main

import (
	"fmt"
	"os"
	"strconv"
)

// QuadB function takes two integers x and y and prints a pattern based on their values
func QuadB(x, y int) {
	// Check if x and y are both greater than 0
	if x > 0 && y > 0 {
		// Print the top row pattern
		fmt.Print("/")
		for i := 0; i < x-2; i++ {
			fmt.Print("*")
		}
		// Print the last character of the top row
		if x > 1 {
			fmt.Println("\\")
		} else {
			fmt.Println()
		}

		// Loop to print the middle rows pattern
		for j := 0; j < y-2; j++ {
			// Print the leftmost character of each row
			fmt.Print("*")
			// Print spaces in between (if any)
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
			}
			// Print the rightmost character of each row
			if x > 1 {
				fmt.Println("*")
			} else {
				fmt.Println()
			}
		}

		// Print the bottom row pattern if y is greater than 1
		if y > 1 {
			// Print the leftmost character of the bottom row
			fmt.Print("\\")
			// Print the middle characters of the bottom row
			for i := 0; i < x-2; i++ {
				fmt.Print("*")
			}
			// Print the rightmost character of the bottom row
			if x > 1 {
				fmt.Println("/")
			} else {
				fmt.Println()
			}
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Error: length of Arguments is not 3\n")
		return
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil || width <= 0 { // Added condition to check for negative or zero width
		fmt.Println("Not a quad function")
		return
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil || height <= 0 { // Added condition to check for negative or zero height
		fmt.Println("Not a quad function")
		return
	}
	QuadB(width, height)
}

