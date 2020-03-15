package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mbauhardt/moneyflow/cmd"
)

func contains(tags []string, tag string) bool {
	for _, a := range tags {
		if a == tag {
			return true
		}
	}
	return false
}

func main() {
	argsWithoutProg := os.Args[1:]
	filter := func(folder string) {
		file, err := os.Open(folder + "/tags")
		cmd.Check(err)
		defer file.Close()

		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			tag := fileScanner.Text()
			if contains(argsWithoutProg, tag) {
				fmt.Println(folder)
				return
			}
		}
	}
	cmd.ScanStdin(filter)
}
