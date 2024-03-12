package files

import (
	"fmt"
	"os"
)

type TempDir struct {
	Name string
}

func NewTempDir(name string) TempDir{
	return TempDir{
		Name: name,
	}
}

func (t *TempDir) GetTempDir() bool {
	_, err := os.Stat(t.Name)
	
	if err != nil {
		return false
	}
	
	return true
}

func (t *TempDir) RemoveTempDir() {
	if err := os.RemoveAll(t.Name); err != nil {
		fmt.Printf("latexresume temporal directory couldn't be removed\n")
		os.Exit(0)
	}
}


func (t *TempDir) CreateTempDir() {
	if err := os.Mkdir(t.Name, os.ModePerm); err != nil {
		fmt.Printf("\nUnable to create temporal directory\n\n")
		os.Exit(0)
	}
}
