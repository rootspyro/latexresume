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
		BuildBasics(resumeData),
		BuildWork(resumeData),	
	)
}

func BuildDocument(header, work string) {
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
%s
\end{document}
		`,
		header,
		work,
	)

	fmt.Println(LaTeXCode)
}

// This function generate the LaTeX code of the Basics Section
func BuildBasics(data JsonResume) string {

	var str string

	// 
	
	str += fmt.Sprintf(
		`
\begin{flushleft}
    \textbf{\LARGE %s} \\
    %s \\
    \href{%s}{%s} | %s | %s \\
    2712 Broadway St, San Francisco, California, CA 94115, US
\end{flushleft}
\section*{Summary}
%s\\
		`,
		data.Basics.Name,
		data.Basics.Label,
		data.Basics.Url,
		data.Basics.Url,
		data.Basics.Email,
		data.Basics.Phone,
		data.Basics.Summary,
	)

	str += `
	`
	return str
} 

func BuildWork(data JsonResume) string {

	var str string

	if len(data.Work) > 0 {

		str = `
\section*{Work Experience}
\begin{itemize}[leftmargin=*]
		`

		for _, work := range data.Work {
			str += fmt.Sprintf(
				`
  \item \entry{%s}{%s} \hfill %s - %s \\
	%s \\
	Highlights:
	\begin{itemize}
				`,
				work.Position,
				work.CompanyName,
				work.StartDate,
				work.EndDate,
				work.Summary,
			)

			// Print the list of highlights
			for _, highlight := range work.Highlights {
				str += fmt.Sprintf(
					`
		\item %s				
					`,
					highlight,
				)
			}

			str += `
	\end{itemize}
			`
		}

		// end of the section
		str += `
\end{itemize}
		`
		
	}

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
