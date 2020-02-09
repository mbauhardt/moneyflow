package cmd

import (
        "os"
        "fmt"
        "bufio"
)

func Check(e error) {
        if e != nil {
                panic(e)
        }
}

func ScanStdin(fns ...func(folder string)) {
        stdinScanner := bufio.NewScanner(os.Stdin)
        for stdinScanner.Scan() {
                fmt.Println()
                folder := stdinScanner.Text()
                for _, fn := range fns {
                        fn(folder)
                }
        }
}
