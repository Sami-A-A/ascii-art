package main

import (
	"ascii/pkg/ascii"
	// "fmt"
	"os"
)

func main() {

	// if len(os.Args) < 2 {
	// 	fmt.Println("ERROR: No arguments to print")
	// 	return
	// } else if len(os.Args) != 2 {
	// 	fmt.Println("ERROR: Place all text in a single argument")
	// 	return
	// } else if len(os.Args[1]) < 1 {
	// 	return
	// } else if os.Args[1] == "\\n" {
	// 	fmt.Println()
	// 	return
	// }

	input := os.Args[1]

	ascii.printAsciiArt(input)
}
