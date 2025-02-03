package main

import (
	"fmt"
	"github.com/Kaya-Sem/oopsie"
)

func main() {
	// basic example

	fmt.Println(
		oopsie.CreateOopsie().Title("Database connection failed.").IndicatorMessage("SQL").Render())

	// Create an empty oopsie and customize it
	customOops2 := oopsie.CreateOopsie()
	customOops2.IndicatorColors(oopsie.BLACK, oopsie.GREEN)
	customOops2.DisableIndicator(false)
	customOops2.IndicatorMessage("200 OK")
	customOops2.Title("connection succesfull")
	fmt.Println(customOops2.Render()) // Output: CRITICAL: Database Connection Failed

	fmt.Println()

	// Create a default oopsie
	defaultOops := oopsie.CreateOopsie()
	defaultOops.IndicatorColors(oopsie.BLACK, oopsie.BRIGHT_RED)
	fmt.Println(defaultOops.Render()) // Output: error: an error occurred

	fmt.Println(oopsie.CreateOopsie().IndicatorColors(oopsie.BRIGHT_GREEN, oopsie.BLACK).IndicatorMessage("INFO").Title("Handshake initialised. Awaiting response from upstream service").Error(fmt.Errorf("Transaction ID: 0x7F3A9B")).Render())

	fmt.Println(oopsie.CreateOopsie().IndicatorColors(oopsie.WHITE, oopsie.BLUE).IndicatorMessage("INFO").Title("Handshake initialised. Awaiting response from upstream service").Error(fmt.Errorf("Transaction ID: 0x7F3A9B")).Render())

	fmt.Println(oopsie.CreateOopsie().IndicatorColors(oopsie.BRIGHT_WHITE, oopsie.RED).IndicatorMessage("504").Title("API Timeout").Error(fmt.Errorf("Connection timeout occured. Perhaps your time/date is too far into the future?")).Render())

}
