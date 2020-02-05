package main

import (
	"os"
	"fmt"
	"bufio"
	"path/filepath"
	"github.com/fatih/color"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mfprint(c *color.Color, f string, pre string) {
	file, err := os.Open(f)
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	c.Print(pre)
	b := false
	for fileScanner.Scan() {
		if (b) {
			c.Print(", ")
		}
		b = true
    		c.Print(fileScanner.Text())
	}
}

func main() {
	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		fmt.Println()
		folder := stdinScanner.Text()
		mfprint(color.New(color.FgCyan), folder + "/description", filepath.Base(folder) + ". ")
		fmt.Println()
		mfprint(color.New(color.FgYellow), folder + "/money", "€ ")
		fmt.Println()
		mfprint(color.New(color.FgBlue), folder + "/tags", "# ")
		fmt.Println()
	}
}
