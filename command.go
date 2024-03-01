/*

	MODULE: COMMAND
	DESCRIPTION: This module has the charge of read the command line inputs

*/

package main

import (
	"flag"
)

type Flags struct {
	InputPath  string
	OutputPath string
}

type Command struct {
	Flags Flags
}

func NewCommand() Command {
	var input *string
	var output *string

	input = flag.String("f", "resume.json", "Specify the input .json file")
	output = flag.String("o", "resume", "Specify the output filename")

	flag.Parse()

	return Command{
		Flags: Flags{
			InputPath:  *input,
			OutputPath: *output,
		},
	}
}

