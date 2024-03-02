/*

	MODULE: COMMAND
	DESCRIPTION: This module has the charge of read the command line inputs

*/

package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	InputPath  string
	OutputPath string
	TEX        bool // Specify to only generate .tex file
	PDF        bool // Specify to only generate .pdf file
	Version    bool
}

type Command struct {
	Flags   Flags
	Version string
}

func NewCommand() Command {
	var input *string
	var output *string
	var version *bool
	var tex *bool
	var pdf *bool

	input = flag.String("json", "resume.json", "Specify the input .json file")
	output = flag.String("o", "resume", "Specify the output filename")
	version = flag.Bool("v", false, "Display the latexresume version")
	tex = flag.Bool("tex", false, "Generate only the .tex result")
	pdf = flag.Bool("pdf", false, "Generate only the .pdf result")

	flag.Parse()

	return Command{
		Version: "1.0",
		Flags: Flags{
			InputPath:  *input,
			OutputPath: *output,
			TEX:        *tex,
			PDF:        *pdf,
			Version:    *version,
		},
	}
}

func (c *Command) PrintVersion() {
	fmt.Printf("latexresume v.%s\n", c.Version)
}
