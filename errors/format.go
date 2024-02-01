package errorhandler

import (
	"strings"
)

// func checkOptionIsValid(arg string) error {
// 	var err error
// 	if strings.HasPrefix(arg, "--"){
// 		switch {
// 		case strings.HasPrefix(arg, "--align="):
// 			alignFlagCheck = true
// 		case strings.HasPrefix(arg, "--output="):
// 			outputFlagCheck = true
// 		case strings.HasPrefix(arg, "--color="):
// 			colorFlagCheck = true
// 		default:
// 			return err
// 		}
// 	} else {
// 		optionCheck = false
// 		stringCheck = true
// 	}
// 	return err
// }


func CheckIsValidFlag(arg string) bool {

	parsedFlags := []string{"align", "color", "output", "reverse"}

	for _, parsedFlag := range parsedFlags {
		if strings.HasPrefix(arg, "--" + parsedFlag + "=") {
			return true
		}
	}

	return false
}

func CheckIsValidBanner(arg string) bool {

	banners := []string{"standard", "tinkertoy", "shadow"}

	for _, banner := range banners {
		if arg == banner {
			return true
		}
	}
	return false
}