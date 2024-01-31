package errorhandler

// import (
// 	"errors"
// 	"strings"
// )



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


// func checkIsValidFlag(arg string) bool {
// 	parsedFlags := []string{"align", "color", "output"}
// 	for _, parsedFlag := range parsedFlags {
// 		if strings.HasPrefix(arg, "--" + parsedFlag + "=") {
// 			return true
// 		}
// 	}
// 	return false
// }