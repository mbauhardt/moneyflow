package persistence

import "os"

type Environment struct {
  DbPath string;
}

func Env() (*Environment, error) {
  hdir, e := os.UserHomeDir()
  if e != nil {
    return nil, e
  }
  return &Environment { DbPath: hdir + "/.local/share/moneyflow/db" }, nil
}
