package main

import (
	"fmt"
	"github.com/Kaya-Sem/oopsie"
)

func main() {
	// Create an empty oopsie and customize it
	customOops := oopsie.CreateEmptyOopsie()
	customOops.SetTitle("Database Connection Failed")
	customOops.SetIndicatorContent("CRITICAL")
	fmt.Println(customOops.Render()) // Output: CRITICAL: Database Connection Failed

	// Create a default oopsie
	defaultOops := oopsie.CreateOopsie()
	fmt.Println(defaultOops.Render()) // Output: error: an error occurred

	// Disable the indicator
	customOops.DisableIndicator()
	fmt.Println(customOops.Render()) // Output: Database Connection Failed
}
