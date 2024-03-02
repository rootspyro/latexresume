package main

import "os"

func main() {

	// Create the command object
	command := NewCommand()

	if command.Flags.Version {
		command.PrintVersion()
		os.Exit(0)
	}

	inputPath := command.Flags.InputPath
	outputPath := command.Flags.OutputPath

	// Get the json schema
	files := NewFiles(inputPath, outputPath)
	resumeData := files.GetJsonData()

	// Build the LaTeX
	latex := NewLatex(resumeData)

	// Write the document
	files.WriteTex(latex.LatexCode)
	files.MakePDF()
}
