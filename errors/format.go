package errorhandler

import (
	"strings"
)

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