package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var hillaryscriptpp = false

func main() {
	red := color.New(color.FgRed)
	// boldRed := red.Add(color.Bold)

	blue := color.New(color.FgBlue)
	// boldBlue := blue.Add(color.Bold)

	yellow := color.New(color.FgYellow)
	// boldYellow := yellow.Add(color.Bold)

	color.Red("Welcome to TrumpScript++!")

	fmt.Println()

	blue.Printf("To switch to HillaryScript++ mode, type ")
	yellow.Printf("EMAILS EMAILS EMAILS")
	blue.Println(" in full caps.")

	red.Printf("To continue in TrumpScript++ mode, type ")
	yellow.Printf("PRESIDENT OF THE VIRGIN ISLANDS")
	red.Println(" in full caps.")

	fmt.Println()

	isTrumpSupporter()
}

func isTrumpSupporter() bool {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	blue := color.New(color.FgBlue)
	boldBlue := blue.Add(color.Bold)

	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)

	statusConfirmed := false
	for !statusConfirmed {
		answer := ""
		prompt := &survey.Input{
			Message: "Of which mode do you want to enter?",
		}
		survey.AskOne(prompt, &answer, nil)

		if answer == "EMAILS EMAILS EMAILS" {
			boldBlue.Println("ALL HAIL OUR LOW IQ GENIUS HILLARY CLINTON AND HER PRIVATE EMAIL SERVERS!")
			time.Sleep(time.Second * 2)
			statusConfirmed = true
		} else if answer == "PRESIDENT OF THE VIRGIN ISLANDS" {
			boldRed.Println("ALL HAIL OUR 70 YEAR OLD PRESIDENT IN ACCELERATED READING!")
			time.Sleep(time.Second * 2)
			statusConfirmed = true
		} else {
			boldGreen.Println("THAT IS NOT AN OPTION! SCREW OFF COMMUNIST PIGS!")
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}
	}
	return false
}

func clear() {

	var c *exec.Cmd
	var doClear = true

	switch runtime.GOOS {
	case "darwin":
	case "linux":
		c = exec.Command("clear")
	case "windows":
		c = exec.Command("cmd", "/c", "cls")
	default:
		doClear = false
	}
	if doClear {
		c.Stdout = os.Stdout
		c.Run()
	}
}
