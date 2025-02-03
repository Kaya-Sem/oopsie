package main

import (
	"fmt"
	"github.com/Kaya-Sem/oopsie"
)

func main() {
	// Create an empty oopsie and customize it
	customOops := oopsie.CreateOopsie()
	customOops.SetTitle("Database Connection Failed")
	customOops.SetIndicatorMessage("CRITICAL")
	fmt.Println(customOops.Render()) // Output: CRITICAL: Database Connection Failed

	fmt.Println()

	// Create an empty oopsie and customize it
	customOops2 := oopsie.CreateOopsie()
	customOops2.SetIndicatorColors(oopsie.BLACK, oopsie.GREEN)
	customOops2.DisableIndicator(false)
	customOops2.SetIndicatorMessage("200 OK")
	customOops2.SetTitle("connection succesfull")
	fmt.Println(customOops2.Render()) // Output: CRITICAL: Database Connection Failed

	fmt.Println()

	// Create a default oopsie
	defaultOops := oopsie.CreateOopsie()
	defaultOops.SetIndicatorColors(oopsie.BLACK, oopsie.BRIGHT_RED)
	fmt.Println(defaultOops.Render()) // Output: error: an error occurred

	// Disable the indicator
	customOops.DisableIndicator(false)
	fmt.Println(customOops.Render()) // Output: Database Connection Failed
}
