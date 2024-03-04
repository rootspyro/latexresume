/*

	MODULE: LATEX 
	DESCRIPTION: This module is focuses on maintaining the LaTeX code

*/

package main

import "fmt"

type LaTeX struct {
	ResumeSchema JsonResume	
	LatexCode string
}

func NewLatex(resume JsonResume) LaTeX {

	var latex LaTeX

	latex.ResumeSchema = resume
	latex.BuildDocument()

	return latex
}

func(l *LaTeX) BuildDocument() {
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
		l.buildBasics(l.ResumeSchema.Basics),
		l.buildWork(l.ResumeSchema.Work),
		l.buildVolunteer(l.ResumeSchema.Volunteer),
		l.buildProjects(l.ResumeSchema.Projects),
		l.buildEducation(l.ResumeSchema.Education),
		l.buildAwards(l.ResumeSchema.Awards),
		l.buildCertificates(l.ResumeSchema.Certificates),
		l.buildPublications(l.ResumeSchema.Publications),
		l.buildSkills(l.ResumeSchema.Skills),
		l.buildAwards(l.ResumeSchema.Awards),
		l.buildInterest(l.ResumeSchema.Interests),
		l.buildReferences(l.ResumeSchema.References),
	)

	l.LatexCode = LaTeXCode
}

func(l *LaTeX) buildBasics(data Basics) string {
	var str string

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
		data.Name,
		data.Label,
		data.Url,
		data.Url,
		data.Email,
		data.Phone,
		data.Location.Address,
		data.Location.City,
		data.Location.Region,
		data.Location.PostalCode,
		data.Location.CountryCode,
		data.Summary,
	)

	return str
}

func(l *LaTeX) buildWork(data []Work) string {

	var str string

	if len(data) > 0 {

		str = `
\section*{Work Experience}
\begin{itemize}[leftmargin=*]
		`

		for _, work := range data {
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

func(l *LaTeX) buildVolunteer(data []Volunteer) string {

	var str string

	if len(data) > 0 {

		str = `
\section*{Volunteer Experience}
\begin{itemize}[leftmargin=*]
		`
		for _, volunteer := range data {
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

func(l *LaTeX) buildEducation(data []Education) string {
	var str string

	if len(data) > 0 {
		str = `
\section*{Education}
\begin{itemize}[leftmargin=*]
		`

		for _, education := range data {
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

func(l *LaTeX) buildAwards( data []Award ) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{Awards}
\begin{itemize}[leftmargin=*]
		`
		for _, award := range data {
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

func(l *LaTeX) buildCertificates(data []Certificate) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{Certificates}
\begin{itemize}[leftmargin=*]
		`
		for _, certificate := range data {
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

func(l *LaTeX) buildPublications(data []Publication) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{Publications}
\begin{itemize}[leftmargin=*]
		`

		for _, publications := range data {

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

func(l *LaTeX) buildSkills(data []Skill) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{Skills}
\begin{itemize}[leftmargin=*]
		`

		for _, skill := range data {
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

func(l *LaTeX) buildLanguages(data []Language) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{Languages}
\begin{itemize}[leftmargin=*]
		` 

		for _, language := range data {
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

func(l *LaTeX) buildProjects(data []Project) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{Projects}
\begin{itemize}[leftmargin=*]
		` 
		
		for _, project := range data {

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

func(l *LaTeX) buildInterest(data []Interest) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{Interests}
\begin{itemize}[leftmargin=*]
		` 

		for _, interest := range data {
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

func(l *LaTeX) buildReferences(data []Reference) string {
	var str string

	if len(data) > 0 {

		str = `
\section*{References}
\begin{itemize}[leftmargin=*]
		` 

		for _, reference := range data {
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

