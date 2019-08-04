package main

import (
	//"fmt"
	"github.com/mbauhardt/moneyflow/persistence"
)

func main() {

	env, err := persistence.Env()
	if err != nil {
		panic(err)
	}
	persistence.NewDatabaseDocument(env)
}
