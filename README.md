# Latexresume

Latexresume is a CLI Tool to automate the generation and update of a professional IT resume.

## System Requirements
Latexresume can convert the data of the JSON schema into LaTeX code without any extra package, but to convert LaTeX code into a formatted PDF resume it is required to install the following packages: 
- Latex
    - To install LaTeX you can install the package [texlive-full](https://www.tug.org/texlive/) or if you don't need the all packages of latex you can install the following:
        - texlive-base
        - texlive-latex-recommend
        - texlive-latex-extra

- Latex Compilation
    - [latexmk](https://mg.readthedocs.io/latexmk.html)

## Installation

### Releases
Check the [releases](https://github.com/rootspyro/latexresume/releases) to download the binary for your OS.

### Build Locally
```shell
git clone https://github.com/rootspyro/latexresume.git
cd latexresume
go build
```

You can validate if it successfully installed running the command `latexresume -v` 
```shell
latexresume -v
# output: latexresume v1.x.x
```

## Usage
latexresume command will read the resume data from a .json file based on the [jsonresume schema](https://jsonresume.org/schema/) and then will convert it in LaTeX code that the user will be able to edit it, also the tool will automatically compile the LaTeX code into a resume in PDF format.

### Basic Usage
```shell
latexresume -json resume.json

# output: 
# resume.tex successfully created! 
# resume.pdf successfully created!
```

### Flags

```shell
Usage of latexresume:
  -json string
    	Specify the input .json file (default "resume.json")
  -o string
    	Specify the output filename (default "resume")
  -pdf
    	Generate only the .pdf result
  -tex
    	Generate only the .tex result
  -v	Display the latexresume version

  -h, --help
        Display the list of flags      
```

## Example

At the [example](example/) directory there is three example file, the input and the two outputs expected of using the `latexresume -json resume.json` command:

- Input: [resume.json](example/resume.json)
- Outout:
    - [resume.tex](example/resume.tex)
    - [resume.pdf](example/resume.pdf)


## License
[MIT](LICENSE)
