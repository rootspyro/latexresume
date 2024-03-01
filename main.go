package main

func main() {
	// Create the command object
	command := NewCommand()
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
