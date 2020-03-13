package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mbauhardt/moneyflow/cmd"
)

func printMoney(folder string) {
	dat, err := ioutil.ReadFile(folder + "/money")
	cmd.Check(err)
	color.Yellow("€ " + string(dat))
}

func newline(folder string) {
	fmt.Println()
}

func main() {
	var in int64
	var out int64
	var diff int64
	calc := func(folder string) {
		dat, err := ioutil.ReadFile(folder + "/money")
		cmd.Check(err)
		byteToInt, err2 := strconv.ParseInt(strings.TrimSpace(string(dat)), 10, 32)
		cmd.Check(err2)
		diff += byteToInt
		if byteToInt > 0 {
			in += byteToInt
		} else if byteToInt < 0 {
			out += byteToInt
		}
	}
	cmd.ScanStdin(calc)
	fmt.Printf("In: €%d, Out: €%d, Diff: €%d\n", in, out, diff)
}
