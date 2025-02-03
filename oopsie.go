package oopsie

import "fmt"

const (
	// Base colors that will be converted to FG/BG variants internally
	BLACK         = "BLACK"
	RED           = "RED"
	GREEN         = "GREEN"
	YELLOW        = "YELLOW"
	BLUE          = "BLUE"
	PURPLE        = "PURPLE"
	CYAN          = "CYAN"
	WHITE         = "WHITE"
	BRIGHT_BLACK  = "BRIGHT_BLACK"
	BRIGHT_RED    = "BRIGHT_RED"
	BRIGHT_GREEN  = "BRIGHT_GREEN"
	BRIGHT_YELLOW = "BRIGHT_YELLOW"
	BRIGHT_BLUE   = "BRIGHT_BLUE"
	BRIGHT_PURPLE = "BRIGHT_PURPLE"
	BRIGHT_CYAN   = "BRIGHT_CYAN"
	BRIGHT_WHITE  = "BRIGHT_WHITE"
)

// Internal color code mapping
var colorCodes = map[string]struct {
	fg string
	bg string
}{
	BLACK:         {fg: "\033[30m", bg: "\033[40m"},
	RED:           {fg: "\033[31m", bg: "\033[41m"},
	GREEN:         {fg: "\033[32m", bg: "\033[42m"},
	YELLOW:        {fg: "\033[33m", bg: "\033[43m"},
	BLUE:          {fg: "\033[34m", bg: "\033[44m"},
	PURPLE:        {fg: "\033[35m", bg: "\033[45m"},
	CYAN:          {fg: "\033[36m", bg: "\033[46m"},
	WHITE:         {fg: "\033[37m", bg: "\033[47m"},
	BRIGHT_BLACK:  {fg: "\033[90m", bg: "\033[100m"},
	BRIGHT_RED:    {fg: "\033[91m", bg: "\033[101m"},
	BRIGHT_GREEN:  {fg: "\033[92m", bg: "\033[102m"},
	BRIGHT_YELLOW: {fg: "\033[93m", bg: "\033[103m"},
	BRIGHT_BLUE:   {fg: "\033[94m", bg: "\033[104m"},
	BRIGHT_PURPLE: {fg: "\033[95m", bg: "\033[105m"},
	BRIGHT_CYAN:   {fg: "\033[96m", bg: "\033[106m"},
	BRIGHT_WHITE:  {fg: "\033[97m", bg: "\033[107m"},
}

const RESET = "\033[0m"

type indicator struct {
	isShown         bool
	content         string
	foregroundColor string
	backgroundColor string
}

func (i *indicator) buildIndicator() string {
	if !i.isShown {
		return ""
	}

	colorCode := ""
	if i.foregroundColor != "" {
		if codes, exists := colorCodes[i.foregroundColor]; exists {
			colorCode += codes.fg
		}
	}
	if i.backgroundColor != "" {
		if codes, exists := colorCodes[i.backgroundColor]; exists {
			colorCode += codes.bg
		}
	}

	return fmt.Sprintf("%s %s %s ", colorCode, i.content, RESET)
}

type Oopsie struct {
	indicator indicator
	title     string
	err       error
}

func (o *Oopsie) Render() string {
	message := fmt.Sprintf("%s%s\n", o.indicator.buildIndicator(), o.title)
	return message
}

func (o *Oopsie) Title(title string) *Oopsie {
	o.title = title

	return o
}

func (o *Oopsie) IndicatorMessage(content string) *Oopsie {
	o.indicator.content = content

	return o
}

func (o *Oopsie) IndicatorColors(fg, bg string) *Oopsie {
	o.indicator.foregroundColor = fg
	o.indicator.backgroundColor = bg

	return o
}

func (o *Oopsie) DisableIndicator(status bool) *Oopsie {
	o.indicator.isShown = !status

	return o
}

func (o *Oopsie) Error(err error) *Oopsie {
	o.err = err

	return o
}

func CreateOopsie() *Oopsie {
	oops := Oopsie{
		indicator: indicator{
			content:         "ERROR",
			isShown:         true,
			foregroundColor: BLACK,
			backgroundColor: RED,
		},
		title: "this is a default error message",
		err:   fmt.Errorf("something went wrong"),
	}
	return &oops
}
