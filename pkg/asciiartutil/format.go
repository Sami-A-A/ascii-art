package asciiartutil

import (
	"fmt"
	"strings"
)

type Color struct {
	AnsiCode string
	LettersToColor string 
}

type AsciiArt struct {
	Text string
	Reverse bool
	OutputFile string
	Alignment string
	Color Color
	Lines [][]string
	Font string
}

func CheckFormat(args []string) (string, string, string, error) {

	// Check Stages
	optionCheck := true
	bannerCheck := false
	colorCheck := false

	// Initializing output variables
	var lettersToColor string
	var textToAscii string
	var banner string
	var err error

	// LOOP THROUGH ALL THE ARGS
	for i, arg := range args {

		if CheckIsValidFlag(arg){
			
			if colorCheck {
				colorCheck = false
			}
			if strings.HasPrefix(arg, "--color=") && !colorCheck {
				colorCheck = true
			}
			if !optionCheck || len(args[i:]) == 1 {
				err = fmt.Errorf("error: formatting error and i = %d", len(args[i:]))
				break
			}

		} else if optionCheck {

			textToAscii = arg
			if colorCheck {
				lettersToColor = arg
				colorCheck = false
			} else {
				optionCheck = false
				bannerCheck = true
			}
			
		} else if bannerCheck {

			if len(args[i:]) > 1 {
				err = fmt.Errorf("error: too many arguments")
			}
			if CheckIsValidBanner(arg) {
				banner = arg
			} else {
				err = fmt.Errorf("error: invalid banner")
			}
		}

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
									// IF NOT    -->    ( MEANS CURRENT ARG IS THE MAIN TEXT AND WE NEED TO CHECK IF NEXT ARG IS BANNER )
										// optionCheck = false
										// bannerCheck = true
										// textToAscii = <current arg>
										// CONTINUE AND CHECK VALID BANNER
											// IF YES    -->    ( MEANS CURRENT ARG IS BANNER AND WE NEED TO CHECK IF ITS THE LAST ARGUMENT )
												// banner = <current arg>
												// CONTINUE AND CHECK IF THERE ARE REMAINING ARGS
													//IF YES    -->    ( MEANS WRONG FORMATTING OR TOO MANY ARGUMENTS AND ERROR IS RETURNED )
														// err = new(error, "Wrong formatting or too many arguments ")
														// return "", "", "", err
											// IF NOT    -->    ( MEANS EVERYTHING IS OK )
					// IF NOT ()
						// CONTINUE UNTIL NOT A FLAG
			// IF NOT    -->    (MEANS WE HAVE NO FLAGS AND CAN GO DIRECTLY TO TEXT)
				// optionCheck = false
				// bannerCheck = true
				// text = <current arg>
				// CONTINUE AND CHECK IF VALID BANNER
					// IF YES
						// banner = <current arg>
						// CONTINUE AND CHECK IF THERE ARE REMAINING ARGS
							// IF YES    -->    ( MEANS WRONG FORMATTING OR TOO MANY ARGUMENTS AND ERROR IS RETURNED )
								// err = new(error, "Wrong formating or too many arguments")
								// return "", "", "", err
							// IF NOT    -->    ( MEANS EVERYTHING IS OK )
					// IF NOT    -->    ( MEANS FORMATTING ERROR )
						// err = new(error, "Wrong formatting or too many arguments")
						// return "", "", "", err
	}

	// Usage: go run . [OPTION] [STRING] [BANNER]
	// EX: go run . --output=<fileName.txt> something standard

	return lettersToColor, textToAscii, banner, err
}

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