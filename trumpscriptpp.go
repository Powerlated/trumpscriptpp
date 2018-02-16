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
	color.Red("Welcome to TrumpScript++!")

	fmt.Printf(color.BlueString("To switch to HillaryScript++ mode, type "))
	fmt.Printf(color.YellowString("EMAILS EMAILS EMAILS"))
	fmt.Println(color.BlueString(" in full caps."))

	fmt.Println()

	fmt.Printf(color.BlueString("To continue with TrumpScript++ mode, type "))
	fmt.Printf(color.YellowString("PRESIDENT OF THE VIRGIN ISLANDS"))
	fmt.Println(color.BlueString(" in full caps."))

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
			statusConfirmed = true
		} else if answer == "PRESIDENT OF THE VIRGIN ISLANDS" {
			boldRed.Println("ALL HAIL OUR 70 YEAR OLD PRESIDENT IN ACCELERATED READING!")
			statusConfirmed = true
		} else {
			boldGreen.Println("THAT IS NOT AN OPTION! SCREW OFF COMMUNIST PIGS!")
			time.Sleep(time.Second * 2)
			clear()
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
