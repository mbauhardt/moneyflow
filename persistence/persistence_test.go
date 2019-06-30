package persistence

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	hdir, _ := os.UserHomeDir()
	tests := []struct {
		name     string
		expected Environment
	}{

		{"XDG Env", Environment{DbPath: hdir + "/.local/share/moneyflow/db"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Env()
			if got.DbPath != tt.expected.DbPath {
				t.Errorf("Env() == %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestIdGenerator(t *testing.T) {
	dir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(dir) // clea

	tests := []struct {
		name  string
		env   Environment
		files []string
		id    string
	}{
		{"no files", Environment{DbPath: dir + "/a"}, []string{}, "0"},
		{"one file 0", Environment{DbPath: dir + "/b"}, []string{"0"}, "1"},
		{"two files 0,1", Environment{DbPath: dir + "/c"}, []string{"0", "1"}, "2"},
		{"one file 1", Environment{DbPath: dir + "/d"}, []string{"1"}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, f := range tt.files {
				p := tt.env.DbPath + "/" + f
				e := os.MkdirAll(p, os.ModePerm)
				fmt.Println(e)
			}
			docid, _ := GenerateDocumentId(&tt.env)
			if docid != tt.id {
				t.Errorf("TestIdGenerator() == %q, want %q", docid, tt.id)
			}
		})
	}

}
