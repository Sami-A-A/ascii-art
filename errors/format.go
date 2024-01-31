package errorhandler

import (
	"errors"
	"fmt"
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


func CheckIsValidFlag(arg string) error {

	parsedFlags := []string{"align", "color", "output", "reverse"}
	var err error

	if !strings.HasPrefix(arg, "--") {
		return err
	}

	for _, parsedFlag := range parsedFlags {
		if strings.HasPrefix(arg, "--" + parsedFlag + "=") {
			err = nil
			break
		} else {
			fmt.Errorf("Error: Invalid flag")
		}
	}

	return err
}