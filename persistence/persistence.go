package persistence

import (
	"errors"
	"github.com/mbauhardt/moneyflow/entities"
	"os"
	"strconv"
)

type Environment struct {
	DbPath string
}

type DatabaseDocument struct {
	Id    string
	money *entities.Money
	desc  *string
	tags  *[]entities.Tag
}

func Env() (*Environment, error) {
	hdir, e := os.UserHomeDir()
	if e != nil {
		return nil, e
	}
	return &Environment{DbPath: hdir + "/.local/share/moneyflow/db"}, nil
}

func GenerateDocumentId(env *Environment) (string, error) {
	for i := 0; i < 10000000; i++ {
		path := env.DbPath + "/" + strconv.Itoa(i)
		if !exists(path) {
			return strconv.Itoa(i), nil
		}
	}
	return "", errors.New("unable to generate database id")
}

func NewDatabaseDocument(env *Environment) (*DatabaseDocument, error) {
	id, err1 := GenerateDocumentId(env)
	if err1 != nil {
		return nil, err1
	}
	err2 := os.MkdirAll(env.DbPath+"/"+id, os.ModePerm)
	if err2 != nil {
		return nil, err2
	}
	_, err3 := os.Create(env.DbPath + "/" + id + "/money")
	if err3 != nil {
		return nil, err3
	}
	_, err4 := os.Create(env.DbPath + "/" + id + "/description")
	if err4 != nil {
		return nil, err4
	}
	_, err5 := os.Create(env.DbPath + "/" + id + "/tags")
	if err5 != nil {
		return nil, err5
	}
	return &DatabaseDocument{Id: id}, nil
}

func SaveTags(env *Environment, doc *DatabaseDocument, tags []entities.Tag) (*DatabaseDocument, error) {
	f, err := os.Create(env.DbPath + "/" + doc.Id + "/tags")
	if err != nil {
    		panic(err)
	}
	defer f.Close()
	for _,t := range tags {
		if t.Modifier == "+" {
			f.WriteString(t.Name)
			f.WriteString("\n")
		}
	}
	f.Sync()
	return &DatabaseDocument{Id: doc.Id}, nil
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
