package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mbauhardt/moneyflow/persistence"
)

func main() {
	env, err := persistence.Env()
	if err != nil {
		panic(err)
	}
	files, err := ioutil.ReadDir(env.DbPath)
	for _, file := range files {
		fmt.Println(env.DbPath + "/" + file.Name())
	}
}
