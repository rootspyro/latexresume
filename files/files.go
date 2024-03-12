/*

	MODULE: FILES
	DESCRIPTION: This module has the charge of read and write the external documents

*/

package files

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rootspyro/latexresume/pipes"
)

type Files struct {
	InputPath      string
	OutputFilename string
	TexFilename    string
	TemporalDir    string
}

func NewFiles(i, o, tempDir string) Files {
	return Files{
		InputPath:      i,
		OutputFilename: o,
		TexFilename:    o + ".tex",
		TemporalDir:    tempDir,
	}
}

// Read the .json schema file
func (f *Files) GetJsonData() pipes.JsonResume {
	var resumeData pipes.JsonResume

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

func (f *Files) BuildOutputs(texCode string, texFlag, pdfFlag bool) {
	f.MakeTexCode(texCode, texFlag, pdfFlag)

	if !texFlag || pdfFlag { // In case the user apply the command "latexresume -pdf -tex" because is dumb like me
		f.MakePDF(texFlag, pdfFlag)
	}
}

// Convert the string LaTeX  code into a .tex output
func (f *Files) MakeTexCode(code string, onlyTex, onlyPDF bool) {
	err := os.WriteFile(f.TexFilename, []byte(code), 0664)
	if err != nil {
		fmt.Printf("\nError creating %s file\n\n", f.TexFilename)
		os.Exit(0)
	}

	if !onlyPDF || onlyTex { // latexresume -pdf -tex case

		splitTexFilename := strings.Split(f.TexFilename, "/") // Get only filename if the output value is a dir path. Example: ./output/resume.pdf
		filename := splitTexFilename[len(splitTexFilename) - 1]

		fmt.Printf("%s successfully created!\n", filename)
	}
}

// Run the latexmk -pdf command
func (f *Files) MakePDF(onlyTex, onlyPDF bool) {

	splitOuputName := strings.Split(f.OutputFilename, "/") // Get only filename if the output value is a dir path. Example: ./output/resume.pdf
	pdfFilename := splitOuputName[len(splitOuputName) - 1] + ".pdf"

	cmd := exec.Command("latexmk", "-output-directory="+f.TemporalDir, "-pdf", f.TexFilename)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Something went wrong creating the %s file.\n", pdfFilename)
		os.Exit(0)
	}

	// Move the generated .pdf file out from the temporal directory
	tempPath := f.TemporalDir + "/" + pdfFilename
	currentPath := f.OutputFilename + ".pdf" 

	if err := os.Rename(tempPath, currentPath); err != nil {
		fmt.Printf("\nSomething went wrong creating the %s file.\n\n", pdfFilename)
		os.Exit(0)
	}

	// if -pdf flag is true then delete .tex file
	if !onlyTex && onlyPDF { // latexresume -pdf -tex case

		if err := os.Remove(f.TexFilename); err != nil {
			fmt.Printf("%s couldn't be removed!\n", f.TexFilename)
			os.Exit(0)
		}

	}

	fmt.Printf("%s successfully created!\n", pdfFilename)
}
