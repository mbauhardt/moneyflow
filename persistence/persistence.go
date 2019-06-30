package persistence

import (
	"errors"
	"os"
	"strconv"
)

type Environment struct {
	DbPath string
}

func Env() (*Environment, error) {
	hdir, e := os.UserHomeDir()
	if e != nil {
		return nil, e
	}
	return &Environment{DbPath: hdir + "/.local/share/moneyflow/db"}, nil
}

func GenerateDocumentId(env *Environment) (string, error) {
	for i := 0; i < 10; i++ {
		path := env.DbPath + "/" + strconv.Itoa(i)
		if !exists(path) {
			return strconv.Itoa(i), nil
		}
	}
	return "", errors.New("unable to generate")
}

// exists returns whether the given file or directory exists
func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
