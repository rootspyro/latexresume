/*

	MODULE: FILES
	DESCRIPTION: This module has the charge of read and write the external documents

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Files struct {
	InputPath      string
	OutputFilename string
}

func NewFiles(i, o string) Files {
	return Files{
		InputPath:      i,
		OutputFilename: o,
	}
}

// Read the .json schema file
func (f *Files) GetJsonData() JsonResume {
	var resumeData JsonResume

	data, err := os.ReadFile(f.InputPath)
	if err != nil {
		fmt.Printf("%s not found!\n", f.InputPath)
	  os.Exit(0)	
	}

	err = json.Unmarshal(data, &resumeData)
	if err != nil {
		fmt.Printf("Not Valid Json Schema\n")
		fmt.Printf("Pls follow the resumejson schema from: https://jsonresume.org/schema/\n")
		os.Exit(0)
	}

	return resumeData
}

// Convert the string LaTeX  code into a .tex output
func (f *Files) WriteTex(code string) {

	texFilename := f.OutputFilename + ".tex"
	err := os.WriteFile(texFilename, []byte(code), 0664)
	if err != nil {
		log.Panic(err.Error())
	}
}
