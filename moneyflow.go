package main

import "fmt"
import "regexp"
import "os"
import "strings"

func main() {
    argsWithoutProg := strings.Join(os.Args[1:], " ")
	tagRegex, _ := regexp.Compile("\\+[a-zA-Z]+")
	moneyRegex, _ := regexp.Compile("money:[0-9\\.]+")
	fmt.Println(tagRegex.FindAllString(argsWithoutProg, -1))
	fmt.Println(moneyRegex.FindAllString(argsWithoutProg, -1))
	fmt.Println(tagRegex.ReplaceAllString(argsWithoutProg, " "))
	fmt.Println(moneyRegex.ReplaceAllString(argsWithoutProg, " "))
	fmt.Println(moneyRegex.ReplaceAllString(tagRegex.ReplaceAllString(argsWithoutProg, " "), " "))
}
