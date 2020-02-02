package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mbauhardt/moneyflow/persistence"
	"github.com/mbauhardt/moneyflow/parse"
)

func main() {

	env, err := persistence.Env()
	if err != nil {
		panic(err)
	}
	argsWithoutProg := os.Args[1:]

	// new db
	doc, dberr := persistence.NewDatabaseDocument(env)
	if dberr != nil {
		panic(dberr)
	}

	// parse
	tags := parse.ParseTags(strings.Join(argsWithoutProg, " "))
	money, merr := parse.ParseMoney(strings.Join(argsWithoutProg, " "))
	if merr != nil {
		panic(merr)
	}
	desc, derr := parse.ParseDescription(strings.Join(argsWithoutProg, " "))
	if derr != nil {
		panic(derr)
	}

	// save
	persistence.SaveDescription(env, doc, desc)
	persistence.SaveTags(env, doc, tags)
	if money != nil {
		persistence.SaveMoney(env, doc, money)
	}
	fmt.Println("Added new doc[" + doc.Id + "]")
}
