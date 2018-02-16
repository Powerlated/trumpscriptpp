package main

import (
	"fmt"

	"github.com/fatih/color"
	survey "gopkg.in/AlecAivazis/survey.v1"
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

	isTrumpSupporter()
}

func isTrumpSupporter() bool {
	statusConfirmed := false
	for !statusConfirmed {
		answer := ""
		prompt := &survey.Input{
			Message: "Of which mode do you want to enter?",
		}
		survey.AskOne(prompt, &answer, nil)
	}
	return false
}
