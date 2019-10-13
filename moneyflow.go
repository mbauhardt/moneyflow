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
	action := os.Args[1]
	if action == "add"  {
		doc, err2 := persistence.NewDatabaseDocument(env)
		if err2 != nil {
			panic(err2)
		}
		tags := parse.ParseTags(strings.Join(argsWithoutProg, " "))
		persistence.SaveTags(env, doc, tags)

		fmt.Println("Added new doc[" + doc.Id + "]"+desc)
	} else {
		fmt.Println("Unknown Command...");
	}

}
