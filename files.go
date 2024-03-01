package main

import (
	"encoding/json"
	"log"
	"os"
)

type Files struct {
	InputPath  string
	OutputPath string
}

func NewFiles(i, o string) Files {

	return Files{
		InputPath: i,
		OutputPath: o,
	}
}

// Read the .json schema file
func(f *Files) GetJsonData() JsonResume {
	var resumeData JsonResume

	data, err := os.ReadFile(f.InputPath)
	if err != nil {
		log.Panic(err.Error())
	}

	err = json.Unmarshal(data, &resumeData)
	if err != nil {
		log.Panic(err.Error())
	}

	return resumeData
}

// Convert the string LaTeX  code into a .tex output
func(f *Files) WriteTex(code string) {
	err := os.WriteFile(f.OutputPath, []byte(code), 0664)

	if err != nil {
		log.Panic(err.Error())
	}
}
