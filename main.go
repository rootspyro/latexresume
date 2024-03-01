package main

func main() {

	jsonPath := "./example/resume.json"
	outputPath := "./resume.tex"

	// Get the json schema
	files := NewFiles(jsonPath, outputPath)
	resumeData := files.GetJsonData()

	// Build the LaTeX
	latex := NewLatex(resumeData)

	// Write the document
	files.WriteTex(latex.LatexCode)
}

