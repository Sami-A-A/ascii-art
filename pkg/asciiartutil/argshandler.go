package asciiartutil

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
