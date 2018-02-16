package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	color.Red("Welcome to TrumpScript++!")
	fmt.Printf(color.BlueString("To switch to HillaryScript++ mode, type "))
	fmt.Printf(color.YellowString("\"EMAILS EMAILS EMAILS\""))
	fmt.Println(color.BlueString(" in full caps."))
}
