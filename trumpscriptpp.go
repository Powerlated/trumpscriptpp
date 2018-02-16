package main

import (
	"fmt"
	"os"
	"strings"

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

		if strings.ToLower(answer) == "quit" {
			os.Exit(0)
		} else if answer == "EMAILS EMAILS EMAILS" {
			fmt.Println(color.BlueString("ALL HAIL OUR LOW IQ GENIUS HILLARY CLINTON AND HER PRIVATE EMAIL SERVERS!"))
			statusConfirmed = true
		} else if answer == "PRESIDENT OF THE VIRGIN ISLANDS" {
			fmt.Println(color.RedString("ALL HAIL OUR 70 YEAR OLD PRESIDENT IN ACCELERATED READING!"))
			statusConfirmed = true
		} else {
			fmt.Println(color.GreenString("THAT IS NOT AN OPTION! SCREW OFF COMMUNIST PIGS!"))
			fmt.Println()
		}
	}
	return false
}
