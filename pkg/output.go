package ascii

import (
	"log"
	"os"
)

// Function to export art to specified file
func Output(art string, outputFlag *string) {
	err := os.WriteFile(*outputFlag, []byte(art), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
