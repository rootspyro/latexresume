package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	jsonPath := "./example/resume.json.json"

	resumeData := GetJsonData(jsonPath)

	latexCode := BuildDocument(
		BuildBasics(resumeData),
		BuildWork(resumeData),
		BuildVolunteer(resumeData),
		BuildProjects(resumeData),
		BuildEducation(resumeData),
		BuildAwards(resumeData),
		BuildCertificates(resumeData),
		BuildPublications(resumeData),
		BuildSkills(resumeData),
		BuildLanguages(resumeData),
		BuildInterest(resumeData),
		BuildReferences(resumeData),
	)

	WriteTex(latexCode)
}

// Convert the string LaTeX  code into a .tex output
func WriteTex(code string) {
	err := os.WriteFile("resume.tex", []byte(code), 0664)

	if err != nil {
		fmt.Println(err)
	}
}

func BuildDocument(header, work, volunteer,projects , education, awards, certificates, publications, skills, languages, interests, references string) string {
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
%s
%s
%s
%s
%s
%s
%s
%s
%s
%s
\end{document}
		`,
		header,
		work,
		volunteer,
		projects,
		education,
		awards,
		certificates,
		publications,
		skills,
		languages,
		interests,
		references,
	)

	return LaTeXCode
}

func BuildProjects(data JsonResume) string {
	var str string

	if len(data.Projects) > 0 {

		str = `
\section*{Projects}
\begin{itemize}[leftmargin=*]
		` 
		
		for _, project := range data.Projects {

			str += fmt.Sprintf(
				`
	\item \entry{%s}{%s - %s} 
	\vspace{0.1in}
	%s \\
	\href{%s}{%s}\\
	Highlights:
	\begin{itemize}
				`,
				project.Name,
				project.StartDate,
				project.EndDate,
				project.Description,
				project.Url,
				project.Url,
			)

			for _, highlight := range project.Highlights {
				str += fmt.Sprintf(
					`
				\item %s
					`,
					highlight,
				)
			}	

			str += `
		\end{itemize}
		\vspace{0.2in}
			`
		}

		str += `
\end{itemize}
		`
	}

	return str
}

func BuildLanguages(data JsonResume) string {
	var str string

	if len(data.Languages) > 0 {

		str = `
\section*{Languages}
\begin{itemize}[leftmargin=*]
		` 

		for _, language := range data.Languages {
			str += fmt.Sprintf(
				`
  \item \textbf{%s} - %s 
				`,
				language.Language,
				language.Fluency,
			)
		}

		str += `
\end{itemize}
		` 
	}

	return str
}

func BuildInterest(data JsonResume) string {
	var str string

	if len(data.Interests) > 0 {

		str = `
\section*{Interests}
\begin{itemize}[leftmargin=*]
		` 

		for _, interest := range data.Interests {
			str += fmt.Sprintf(`\item %s: `, interest.Name)

			for i,keyword := range interest.Keywords {
				str += fmt.Sprintf("%s", keyword)

				if i < len(interest.Keywords) - 1 {
					str += fmt.Sprintf(", ")
				}
			}
		}

		str += `
\end{itemize}
		` 
	}

	return str
}

func BuildReferences(data JsonResume) string {
	var str string

	if len(data.References) > 0 {

		str = `
\section*{References}
\begin{itemize}[leftmargin=*]
		` 

		for _, reference := range data.References {
			str += fmt.Sprintf(
				`
	\item \textbf{%s} \\
	%s	

				`,
				reference.Name,
				reference.Reference,
			)
		}

		str += `
\end{itemize}
		` 
	}

	return str
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
    %s, %s, %s, %s, %s
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
		data.Basics.Location.Address,
		data.Basics.Location.City,
		data.Basics.Location.Region,
		data.Basics.Location.PostalCode,
		data.Basics.Location.CountryCode,
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
  \item \entry{%s}{%s} \hfill %s - %s 

	\vspace{0.1in}

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
	\vspace{0.2in}
			`
		}

		// end of the section
		str += `
\end{itemize}
		`
	}

	return str
}

func BuildVolunteer(data JsonResume) string {
	var str string

	if len(data.Volunteer) > 0 {

		str = `
\section*{Volunteer Experience}
\begin{itemize}[leftmargin=*]
		`
		for _, volunteer := range data.Volunteer {
			str += fmt.Sprintf(
				`
	\item \entry{%s}{%s} \hfill %s - %s \\
	%s \\
	Highlights:
	\begin{itemize}
				`,
				volunteer.Position,
				volunteer.Organization,
				volunteer.StartDate,
				volunteer.EndDate,
				volunteer.Summary,
			)

			for _, highlight := range volunteer.Highlights {
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

		str += `\end{itemize}`
	}

	return str
}

func BuildEducation(data JsonResume) string {
	var str string

	if len(data.Education) > 0 {
		str = `
\section*{Education}
\begin{itemize}[leftmargin=*]
		`

		for _, education := range data.Education {
			str += fmt.Sprintf(
				`
	\item \entry{%s}{%s} \hfill %s - %s \\
				`,
				education.Area,
				education.Institution,
				education.StartDate,
				education.EndDate,
			)

			if len(education.Score) > 0 {
				str += fmt.Sprintf(
					`
		Score: %s\\
					`,
					education.Score,
				)
			}

			if len(education.Courses) > 0 {
					str += `
		Courses:
		\begin{itemize}
					`
				for _, course := range education.Courses {
					str += fmt.Sprintf(
						`
			\item %s
						`,
						course,
					)
				}

				str += `\end{itemize}`
			}

		}

		str += `
	\end{itemize}
		`
	}

	return str
}

func BuildAwards(data JsonResume) string {
	var str string

	if len(data.Awards) > 0 {

		str = `
\section*{Awards}
\begin{itemize}[leftmargin=*]
		`
		for _, award := range data.Awards {
			str += fmt.Sprintf(
				`
	\item \textbf{%s} \hfill %s\\
	%s \\
	%s.
				`,	
				award.Title,
				award.Date,
				award.Awarder,
				award.Summary,
			)
		} 

		str += `
\end{itemize}
		`
	}

	return str
}

func BuildCertificates(data JsonResume) string {
	var str string

	if len(data.Certificates) > 0 {

		str = `
\section*{Certificates}
\begin{itemize}[leftmargin=*]
		`
		for _, certificate := range data.Certificates {
			str += fmt.Sprintf(
				`
	\item \textbf{%s} \hfill %s \\
	Issuer: %s \\
	\href{%s}{%s}
				`,
				certificate.Name,
				certificate.Date,
				certificate.Issuer,
				certificate.Url,
				certificate.Url,
			)
		}

		str += `
\end{itemize}
		`
	}	

	return str
}

func BuildPublications(data JsonResume) string {
	var str string

	if len(data.Publications) > 0 {

		str = `
\section*{Publications}
\begin{itemize}[leftmargin=*]
		`

		for _, publications := range data.Publications {

			str += fmt.Sprintf(
				`
	\item \textbf{%s} \hfill %s \\
	Publisher: %s \\
	\href{%s}{%s} \\
	%s
				`,
				publications.Name,
				publications.ReleaseDate,
				publications.Publisher,
				publications.Url,
				publications.Url,
				publications.Summary,
			)
		}

		str += `
\end{itemize}
		`
	}

	return str
}

func BuildSkills(data JsonResume) string {
	var str string

	if len(data.Skills) > 0 {

		str = `
\section*{Skills}
\begin{itemize}[leftmargin=*]
		`

		for _, skill := range data.Skills {
			str += fmt.Sprintf(
				`
	\item \textbf{%s} - %s \\
				`,
				skill.Name,
				skill.Level,
			)

			str += "Keywords: "
			for i, keyword := range skill.Keywords {
				str += fmt.Sprintf("%s", keyword)

				if i < len(skill.Keywords) - 1 {
					str += ", "
				}
			}
		}
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
