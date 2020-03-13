package cmd

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ScanStdin(fns ...func(folder string)) {
	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		folder := stdinScanner.Text()
		for _, fn := range fns {
			fn(folder)
		}
	}
}
