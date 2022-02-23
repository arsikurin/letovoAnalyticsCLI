package colorlib

type (
	background struct {
		Reset, Black, Red, Green, Yellow, Blue, Magenta, Cyan, Beige, White, Grey, LightRed, LightGreen, LightYellow, LightBlue, LightMagenta, LightCyan string
	}
	foreground struct {
		Reset, Black, Red, Green, Yellow, Blue, Magenta, Cyan, Beige, White, Grey, LightRed, LightGreen, LightYellow, LightBlue, LightMagenta, LightCyan string
	}
	style struct {
		Reset, Bold, Dim, Italic, Underlined, Blink, Inverted, Hidden, Strikethrough, UnderlinedDouble, UnderlinedLower string
	}
)

type controlCharacters struct {
	Backspace, Tab, LineFeed, CarriageReturn string
}

var (
	Style        = style{"\033[0m", "\033[1m", "\033[2m", "\033[3m", "\033[4m", "\033[5m", "\033[7m", "\033[8m", "\033[9m", "\033[21m", "\033[52m"}
	Fg           = foreground{"\033[0m", "\033[30m", "\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m", "\033[37m", "\033[38m", "\033[90m", "\033[91m", "\033[92m", "\033[93m", "\033[94m", "\033[95m", "\033[96m"}
	Bg           = background{"\033[0m", "\033[40m", "\033[41m", "\033[42m", "\033[43m", "\033[44m", "\033[45m", "\033[46m", "\033[47m", "\033[48m", "\033[100m", "\033[101m", "\033[102m", "\033[103m", "\033[104m", "\033[105m", "\033[106m"}
	ControlChars = controlCharacters{"\b", "\t", "\n", "\r"}
)
