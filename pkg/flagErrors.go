package ascii

import (
	"flag"
	"fmt"
)

func FlagArgErrors() bool {
	// Error handling; flag arguments: should be a single arg
	if len(flag.Args()) < 1 {
		fmt.Println("ERROR: No arguments to print")
		return true
	} else if len(flag.Args()) > 1 {
		fmt.Println("ERROR: Place all printable text in a single argument")
		return true
	} else if len(flag.Args()[0]) < 1 {
		return true
	}
	return false
}
