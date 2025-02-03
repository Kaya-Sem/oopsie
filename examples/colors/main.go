package main

import (
	"fmt"
	"github.com/Kaya-Sem/oopsie"
)

func main() {
	whiteFgColors := []string{
		oopsie.BLACK, oopsie.RED, oopsie.GREEN, oopsie.YELLOW, oopsie.BLUE, oopsie.PURPLE, oopsie.CYAN}

	blackFgColors := []string{
		oopsie.BRIGHT_BLACK, oopsie.BRIGHT_RED, oopsie.BRIGHT_GREEN, oopsie.BRIGHT_YELLOW, oopsie.BRIGHT_BLUE, oopsie.BRIGHT_PURPLE, oopsie.BRIGHT_CYAN, oopsie.WHITE, oopsie.BRIGHT_WHITE,
	}

	for _, bgColor := range whiteFgColors {
		fmt.Println(
			oopsie.CreateOopsie().Title(fmt.Sprintf("Background color %s", bgColor)).IndicatorMessage("ERROR").IndicatorColors(oopsie.BRIGHT_WHITE, bgColor).Error(fmt.Errorf("")).Render(),
		)
	}

	for _, bgColor := range blackFgColors {
		fmt.Println(
			oopsie.CreateOopsie().Title(fmt.Sprintf("Background color %s", bgColor)).IndicatorMessage("OOPSIE").Error(fmt.Errorf("")).IndicatorColors(oopsie.BLACK, bgColor).Render(),
		)
	}
}
