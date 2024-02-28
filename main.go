package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	jsonPath := "example/resume.json"

	resumeData := GetJsonData(jsonPath)

	BuildDocument(
		HeaderSection(resumeData),
	)
}

func BuildDocument(header string) {
	var LaTeXCode string

	LaTeXCode = `
\documentclass{article}
\usepackage{helvet}
\renewcommand{\familydefault}{\sfdefault}
\usepackage{hyperref}
\usepackage{enumitem}
\usepackage{geometry}

\geometry{
    a4paper,
    total={170mm,257mm},
    left=20mm,
    top=20mm,
}

\newcommand{\entry}[2]{
    \textbf{#1} \hfill \textit{#2}
}

	`

	LaTeXCode += fmt.Sprintf(
		`
\begin{document}
%s
\end{document}
		`,
		header,
	)

	fmt.Println(LaTeXCode)
}

// This function generate the LaTeX code of the Basics Section
func HeaderSection(data JsonResume) string {
	var str string

	str += fmt.Sprintf(
		`
\begin{flushleft}
    \textbf{\LARGE %s} \\
    %s \\
    \href{https://johndoe.com}{johndoe.com} | john@gmail.com | (912) 555-4321 \\
    2712 Broadway St, San Francisco, California, CA 94115, US
\end{flushleft}
		`,
		data.Basics.Name,
		data.Basics.Label,
	)
	return str
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
