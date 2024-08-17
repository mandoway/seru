package syntactic_reducers

import (
	"os/exec"
)

type Functions struct {
	BuildReductionCommand BuildReductionCommandType
	GetOutputFilename     GetOutputFilename
}

type BuildReductionCommandType func(inputFilePath, testFilePath, language string) *exec.Cmd

type GetOutputFilename func(inputFilePath string) string
