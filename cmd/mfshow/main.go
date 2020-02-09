package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
	"path/filepath"
	"github.com/fatih/color"
	"github.com/mbauhardt/moneyflow/cmd"
)

func printDescription(folder string) {
	dat, err := ioutil.ReadFile(folder + "/description")
    	cmd.Check(err)
    	color.Cyan(filepath.Base(folder) + ". " + string(dat))
}

func printMoney(folder string) {
	dat, err := ioutil.ReadFile(folder + "/money")
    	cmd.Check(err)
    	color.Yellow("€ " + string(dat))
}

func printTags(folder string) {
	file, err := os.Open(folder + "/tags")
	cmd.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	c:= color.New(color.FgBlue)
	c.Print("# ")
	b := false
	for fileScanner.Scan() {
		if (b) {
			c.Print(", ")
		}
		b = true
    		c.Print(fileScanner.Text())
	}
	fmt.Println()
}

func newline(folder string) {
	fmt.Println()
}

func main() {
	cmd.ScanStdin(printDescription, printMoney, printTags, newline)
}
