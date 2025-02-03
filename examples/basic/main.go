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

}
