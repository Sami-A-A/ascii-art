package options

func SetFont(input []string) string {
	if input[len(input)-1] == "standard" {
		return ""
	} else {
		return "standard"
	}
}