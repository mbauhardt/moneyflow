package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mbauhardt/moneyflow/cmd"
)

func printTags(folder string) {
	file, err := os.Open(folder + "/tags")
	cmd.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
}

func main() {
	cmd.ScanStdin(printTags)
}
