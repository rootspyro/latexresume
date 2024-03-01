/*

	MODULE: FILES
	DESCRIPTION: This module has the charge of read and write the external documents

*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Files struct {
	InputPath      string
	OutputFilename string
}

func NewFiles(i, o string) Files {
	return Files{
		InputPath:      i,
		OutputFilename: o + ".tex",
	}
}

// Read the .json schema file
func (f *Files) GetJsonData() JsonResume {
	var resumeData JsonResume

	data, err := os.ReadFile(f.InputPath)
	if err != nil {
		fmt.Printf("\n%s file not found!\n", f.InputPath)
		fmt.Printf("Pls follow the resumejson schema from: https://jsonresume.org/schema/\n\n")
		os.Exit(0)
	}

	err = json.Unmarshal(data, &resumeData)
	if err != nil {
		fmt.Printf("\nNot Valid Json Schema\n")
		fmt.Printf("Pls follow the resumejson schema from: https://jsonresume.org/schema/\n\n")
		os.Exit(0)
	}

	return resumeData
}

// Convert the string LaTeX  code into a .tex output
func (f *Files) WriteTex(code string) {
	err := os.WriteFile(f.OutputFilename, []byte(code), 0664)
	if err != nil {
		fmt.Printf("Error creating %s file", f.OutputFilename)
		os.Exit(0)
	}
}

// Run the latexmk -pdf command
func (f *Files) MakePDF() {
	cmd := exec.Command("latexmk", "-pdf", f.OutputFilename)

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(0)
	}
}
