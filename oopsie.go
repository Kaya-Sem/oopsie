package oopsie

import "fmt"

const (
	// Base colors that will be converted to FG/BG variants internally
	BLACK  = "BLACK"
	RED    = "RED"
	GREEN  = "GREEN"
	YELLOW = "YELLOW"
	BLUE   = "BLUE"
	PURPLE = "PURPLE"
	CYAN   = "CYAN"
	WHITE  = "WHITE"
)

// Internal color code mapping
var colorCodes = map[string]struct {
	fg string
	bg string
}{
	BLACK:  {fg: "\033[30m", bg: "\033[40m"},
	RED:    {fg: "\033[31m", bg: "\033[41m"},
	GREEN:  {fg: "\033[32m", bg: "\033[42m"},
	YELLOW: {fg: "\033[33m", bg: "\033[43m"},
	BLUE:   {fg: "\033[34m", bg: "\033[44m"},
	PURPLE: {fg: "\033[35m", bg: "\033[45m"},
	CYAN:   {fg: "\033[36m", bg: "\033[46m"},
	WHITE:  {fg: "\033[37m", bg: "\033[47m"},
}

const RESET = "\033[0m"

type indicator struct {
	isShown         bool
	content         string
	foregroundColor string // Now stores simple color name (e.g., "RED")
	backgroundColor string // Now stores simple color name (e.g., "WHITE")
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

	return fmt.Sprintf("%s%s %s", colorCode, i.content, RESET)
}

type Oopsie struct {
	indicator indicator
	title     string
	err       error
}

func (o *Oopsie) Render() string {
	message := fmt.Sprintf("%s%s", o.indicator.buildIndicator(), o.title)
	return message
}

func (o *Oopsie) SetTitle(title string) {
	o.title = title
}

func (o *Oopsie) SetIndicatorContent(content string) {
	o.indicator.content = content
}

func (o *Oopsie) SetIndicatorColors(fg, bg string) {
	o.indicator.foregroundColor = fg
	o.indicator.backgroundColor = bg
}

func (o *Oopsie) DisableIndicator() {
	o.indicator.isShown = false
}

func CreateEmptyOopsie() *Oopsie {
	oops := new(Oopsie)
	oops.indicator.isShown = true
	return oops
}

func CreateOopsie() *Oopsie {
	oops := Oopsie{
		indicator: indicator{
			content:         "ERROR",
			isShown:         true,
			foregroundColor: WHITE, // Now using simple color name
			backgroundColor: RED,   // Now using simple color name
		},
		title: "an error occurred",
		err:   fmt.Errorf("something went wrong"),
	}
	return &oops
}

func (o *Oopsie) Throw() {
	// Implementation here
}

func throwOopsie(title string, err error) {
	// Implementation here
}
