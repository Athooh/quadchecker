package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str := readInput(reader)
	x, y := countDimensions(str)

	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		return
	}

	checkAndPrintQuad(str, x, y)
}

func readInput(reader *bufio.Reader) string {
	var builder strings.Builder
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		builder.WriteRune(char)
	}
	return builder.String()
}

func countDimensions(str string) (int, int) {
	x := strings.Count(str, "\n")
	y := strings.Index(str, "\n")
	return y, x
}

func checkAndPrintQuad(str string, x, y int) {
	quads := map[string][]rune{
		"quadA": {'o', 'o', 'o', 'o', '-', '|'},
		"quadB": {'/', '\\', '\\', '/', '*', '*'},
		"quadC": {'A', 'A', 'C', 'C', 'B', 'B'},
		"quadD": {'A', 'C', 'A', 'C', 'B', 'B'},
		"quadE": {'A', 'C', 'C', 'A', 'B', 'B'},
	}

	var found bool
	for _, quadName := range []string{"quadA", "quadB", "quadC", "quadD", "quadE"} {
		runes := quads[quadName]
		if matchesQuad(str, x, y, runes...) {
			if found {
				fmt.Print(" || ")
			}
			fmt.Printf("[%s] [%v] [%v]", quadName, x, y)
			found = true
		}
	}

	if found {
		fmt.Println()
		return
	}

	fmt.Println("Not a quad function")
}

func matchesQuad(str string, x, y int, runes ...rune) bool {
	var builder strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0:
				if j == 0 {
					builder.WriteRune(runes[0])
				} else if j == x-1 {
					builder.WriteRune(runes[1])
				} else {
					builder.WriteRune(runes[4])
				}
			case i == y-1:
				if j == 0 {
					builder.WriteRune(runes[2])
				} else if j == x-1 {
					builder.WriteRune(runes[3])
				} else {
					builder.WriteRune(runes[4])
				}
			default:
				if j == 0 || j == x-1 {
					builder.WriteRune(runes[5])
				} else {
					builder.WriteRune(' ')
				}
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String() == str
}
