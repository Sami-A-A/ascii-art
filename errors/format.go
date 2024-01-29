package errorhandler

import (
	// "errors"
	"strings"
)

// Check Stages
var optionCheck = true
var stringCheck = false
var bannerCheck = false

// Check Flags
var colorFlagCheck = false
var outputFlagCheck = false
var alignFlagCheck = false

func CheckFormat(input []string) error {

	// Usage: go run . [OPTION] [STRING] [BANNER]
	// EX: go run . --output=<fileName.txt> something standard
	
	// if len(input) == 0 {
	// 	return errors.New("No argument found")
	// } else if len(input) >= 2 {
	// 	// For each argument check if it's correctly ordered
	// 	for i := 0; i < len(input); i++ {

	// 		if optionCheck {
	// 			checkOptionIsValid(input[i]) 
	// 		}

	// 		if stringCheck {
	// 			if colorFlagCheck {
	// 				colorFlagCheck = false
	// 				optionCheck = true
	// 				stringCheck = 
	// 			}
	// 		}

	// 		if bannerCheck {

	// 		}

	// 	}
	// }
	return nil
}

func checkOptionIsValid(arg string) bool {
	if strings.HasPrefix(arg, "--"){ 
		switch {
		case strings.HasPrefix(arg, "--align="):
			alignFlagCheck = true
		case strings.HasPrefix(arg, "--output="):
			outputFlagCheck = true
		case strings.HasPrefix(arg, "--color="):
			colorFlagCheck = true
		default: 
			return false
		}
	} else {
		optionCheck = false
		stringCheck = true
	}
}


func checkIsValidFlag(arg string) bool {
	parsedFlags := []string{"align", "color", "output"}
	for _, parsedFlag := range parsedFlags {
		if strings.HasPrefix(arg, "--" + parsedFlag + "=") {
			return true
		}
	}
	return false
}