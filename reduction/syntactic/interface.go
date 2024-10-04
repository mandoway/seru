package syntactic

import (
	"github.com/mandoway/seru/reduction/candidate"
	"os/exec"
)

type Functions struct {
	BuildReductionCommand BuildReductionCommandType
	GetOutputFilename     GetOutputFilename
	Init                  InitSyntacticReducer
}

type BuildReductionCommandType func(candidate candidate.Candidate, language string) *exec.Cmd

type GetOutputFilename func(inputFilePath string) string

type InitSyntacticReducer func(language string) error
