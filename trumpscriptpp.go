package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

var red = color.New(color.FgRed)
var boldRed = red.Add(color.Bold)

var blue = color.New(color.FgBlue)
var boldBlue = blue.Add(color.Bold)

var green = color.New(color.FgGreen)
var boldGreen = green.Add(color.Bold)

var yellow = color.New(color.FgYellow)
var boldYellow = yellow.Add(color.Bold)

var loadedFile = ""
var trumpscriptpp = false

func main() {
	color.Red("Welcome to TrumpScript++!")

	fmt.Println()

	blue.Printf("To switch to HillaryScript++ mode, type ")
	boldYellow.Printf("EMAILS EMAILS EMAILS")
	blue.Printf(" or ")
	boldYellow.Printf("1")
	blue.Println(" in full caps.")

	red.Printf("To continue in TrumpScript++ mode, type ")
	boldYellow.Printf("PRESIDENT OF THE VIRGIN ISLANDS")
	red.Printf(" or ")
	boldYellow.Printf("2")
	red.Println(" in full caps.")

	trumpscriptpp = isTrumpSupporter()

	loadFile()
}

func isTrumpSupporter() bool {
	statusConfirmed := false
	for !statusConfirmed {
		fmt.Println()
		answer := ""
		prompt := &survey.Input{
			Message: "Of which mode do you want to enter?",
		}
		survey.AskOne(prompt, &answer, nil)

		if answer == "EMAILS EMAILS EMAILS" || answer == "1" {
			boldBlue.Println("ALL HAIL OUR LOW IQ GENIUS HILLARY CLINTON AND HER PRIVATE EMAIL SERVERS!")
			time.Sleep(time.Second * 2)
			statusConfirmed = true
			return false
		} else if answer == "PRESIDENT OF THE VIRGIN ISLANDS" || answer == "2" {
			boldRed.Println("ALL HAIL OUR 70 YEAR OLD PRESIDENT IN ACCELERATED READING!")
			time.Sleep(time.Second * 2)
			statusConfirmed = true
			return true
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

func loadFile() {
	fileConfirmed := false
	for !fileConfirmed {
		fmt.Println()
		answer := ""
		prompt := &survey.Input{
			Message: "Open a file to run:",
		}
		survey.AskOne(prompt, &answer, nil)
		if strings.ToLower(answer) == "quit" {
			os.Exit(0)
		}
		b, err := ioutil.ReadFile(answer)
		if err != nil {
			if trumpscriptpp {
				red.Println("That file name is Fake News!")
			} else {
				blue.Println("That file does not exist, just like my emails.")
			}
		} else {
			loadedFile = string(b)
			if checkHeader() == false {
				red.Print("This file is Fake News! It does not contain ")
				boldYellow.Print("Make programming great again")
				red.Println(" as a header!")
				os.Exit(0)
			} else if checkFooter() == false {
				red.Print("This file is Fake News! It does not contain ")
				boldYellow.Print("America is great")
				red.Println(" as a footer!")
				os.Exit(0)
			} else {
				fileConfirmed = true
			}
		}
	}
}

func checkHeader() bool {
	if utf8.RuneCountInString(loadedFile) <= 28 {
		return false
	}
	if loadedFile[0:28] == "Make programming great again" {
		return true
	}
	return false
}

func checkFooter() bool {
	if utf8.RuneCountInString(loadedFile) <= 16 {
		return false
	}
	if loadedFile[len(loadedFile)-16:] == "America is great" {
		return true
	}
	return false
}
