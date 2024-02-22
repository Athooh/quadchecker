package main

import (
	"fmt"
	"os"
	"strconv"
)

// QuadE function takes two parameters x and y, representing the number of columns and rows, respectively.
func QuadE(x, y int) {
	// Check if both x and y are greater than 0
	if x > 0 && y > 0 {
		// Print "A" for the first row
		fmt.Print("A")

		// Prints middle elements of the top row by subtracting the width by 2
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}

		// Print "C" if x is greater than 1, otherwise move to the next line
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}

		// Loop through y-2 rows
		for j := 0; j < y-2; j++ {
			// Print "B" for the first column
			fmt.Print("B")

			// Print spaces for x-2 times in each row
			for i := 0; i < x-2; i++ {
				fmt.Print(" ")
			}

			// Print "B" at the end of each row if x is greater than 1, otherwise move to the next line
			if x > 1 {
				fmt.Println("B")
			} else {
				fmt.Println()
			}
		}

		// Print "C" for the last row if y is greater than 1
		if y > 1 {
			// Print "C" for the first column
			fmt.Print("C")

			// Print "B" for x-2 times in the last row
			for i := 0; i < x-2; i++ {
				fmt.Print("B")
			}

			// Print "A" at the end of the last row if x is greater than 1, otherwise move to the next line
			if x > 1 {
				fmt.Println("A")
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
	QuadE(width, height)
}

