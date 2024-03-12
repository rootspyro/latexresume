package main

import (
	"os"

	"github.com/rootspyro/latexresume/files"
)

func main() {
	// Create the command object
	command := NewCommand()

	if command.Flags.Version {
		command.PrintVersion()
		os.Exit(0)
	}

	inputPath := command.Flags.InputPath
	outputPath := command.Flags.OutputPath

	tempDir := files.NewTempDir(".latexresume_temp")

	if !tempDir.GetTempDir() {
		tempDir.CreateTempDir()	
	}

	// always remove the temporal directory at the end of the program
	defer tempDir.RemoveTempDir() 

	// Get the json schema
	f := files.NewFiles(inputPath, outputPath, tempDir.Name)
	resumeData := f.GetJsonData()

	// Build the LaTeX
	latex := NewLatex(resumeData)

	// Write the document
	f.BuildOutputs(
		latex.LatexCode,
		command.Flags.TEX,
		command.Flags.PDF,
	)
}
