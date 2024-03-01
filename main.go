package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	jsonPath := "./example/resume.json"

	resumeData := GetJsonData(jsonPath)

	latex := NewLatex(resumeData)

	WriteTex(latex.LatexCode)
}

// Convert the string LaTeX  code into a .tex output
func WriteTex(code string) {
	err := os.WriteFile("resume.tex", []byte(code), 0664)

	if err != nil {
		fmt.Println(err)
	}
}

func GetJsonData(jsonPath string) JsonResume {
	var resumeData JsonResume

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Panic(err.Error())
	}

	err = json.Unmarshal(data, &resumeData)
	if err != nil {
		log.Panic(err.Error())
	}

	return resumeData
}
