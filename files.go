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
	TexFilename    string
	TemporalDir    string
}

func NewFiles(i, o string) Files {
	return Files{
		InputPath:      i,
		OutputFilename: o,
		TexFilename: o + ".tex",
		TemporalDir: ".latexresume_temp",
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
	err := os.WriteFile(f.TexFilename, []byte(code), 0664)
	if err != nil {
		fmt.Printf("\nError creating %s file\n\n", f.TexFilename)
		os.Exit(0)
	}

	fmt.Printf("%s successfully created!\n", f.TexFilename)
}

// Run the latexmk -pdf command
func (f *Files) MakePDF() {

	pdfFilename := f.OutputFilename + ".pdf"

	// Create temporal directory
	if err := os.Mkdir(f.TemporalDir, os.ModePerm); err != nil {
		fmt.Printf("\nUnable to create temporal directory\n\n")
		os.Exit(0)
	}

	cmd := exec.Command("latexmk", "-output-directory=" + f.TemporalDir, "-pdf", f.TexFilename)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Something went wrong creating the %s file.\n", pdfFilename)
		os.Exit(0)
	}

	// Move the generated .pdf file out from the temporal directory
	tempPath := f.TemporalDir + "/" + pdfFilename
	currentPath := pdfFilename

	if err := os.Rename(tempPath, currentPath); err != nil {
		fmt.Printf("\nSomething went wrong creating the %s.pdf file.\n\n", pdfFilename)
		os.Exit(0)
	}

	// Remove the temporal directory
	if err := os.RemoveAll(f.TemporalDir); err != nil {
		fmt.Println(err)
		fmt.Printf("latexresume temporal directory couldn't be removed\n",)
		os.Exit(0)
	}

	fmt.Printf("%s successfully created!\n", pdfFilename)

}

