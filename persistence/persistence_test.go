package persistence

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	hdir,_ := os.UserHomeDir()
	
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
