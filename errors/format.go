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

func CheckFormat(input []string) (string, string, string, error) {

	var lettersToColor string
	var textToAscii string
	var banner string
	var err error

	// LOOP THROUGH ALL THE ARGS
		// CHECK IF VALID FLAG
			// IF YES    -->    ( MEANS WE'RE STILL LOOKING FOR FLAGS )
				// CHECK IF COLOR FLAG
					// IF YES    -->    ( MEANS WE NEED TO CHECK THE FOLLOWING ARG )
						// CONTINUE AND CHECK IF VALID FLAG
							// IF YES    -->    ( MEANS COLOR HAS NO ARGUMENT )
								// CONTINUE UNTIL NOT A FLAG
							// IF NOT    -->    ( MEANS LETTERS MAY HAVE BEEN PASSED AS ARGUMENT BUT WE NEED TO CHECK IF ITS AN ARGUMENT OR THE MAIN TEXT )
								// lettersToColor = <current arg>
								// CONTINUE AND CHECK IF VALID FLAG
									// IF YES    -->    ( MEANS COLOR HAS ONE ARGUMENT AND WE ARE NOW JUST LOOKING FOR OTHER FLAGS )
										// CONTINUE UNTIL NOT A FLAG
									// IF NOT    -->    ( MEANS CURRENT ARG IS THE MAIN TEXT AND WE NEXT TO CHECK IF NEXT ARG IS BANNER )
										// optionCheck = false
										// bannerCheck = true
										// textToAscii = <current arg>
										// CONTINUE AND CHECK VALID BANNER
											//

					// IF NOT ()
						// CONTINUE UNTIL NOT A FLAG
			// IF NOT    -->    (MEANS WE HAVE NO FLAGS AND CAN GO DIRECTLY TO TEXT)
				// optionCheck = false
				// stringCheck = true









	// Usage: go run . [OPTION] [STRING] [BANNER]
	// EX: go run . --output=<fileName.txt> something standard

	return lettersToColor, textToAscii, banner, err
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