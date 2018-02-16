package main

import (
	"fmt"
	"github.com/fatih/color"
)

var hillaryscriptpp = false

func main() {
	color.Red("Welcome to TrumpScript++!")
	fmt.Printf(color.BlueString("To switch to HillaryScript++ mode, type "))
	fmt.Printf(color.YellowString("\"EMAILS EMAILS EMAILS\""))
	fmt.Println(color.BlueString(" in full caps."))

	fmt.Println()

	fmt.Printf(color.BlueString("To continue with TrumpScript++ mode, type "))
	fmt.Printf(color.YellowString("\"PRESIDENT OF THE VIRGIN ISLANDS\""))
	fmt.Println(color.BlueString(" in full caps."))
}
