package syntactic

import (
	"github.com/mandoway/seru/reduction/domain"
	"os/exec"
)

type Functions struct {
	BuildReductionCommand BuildReductionCommandType
	GetOutputFilename     GetOutputFilename
}

type BuildReductionCommandType func(candidate domain.Candidate, language string) *exec.Cmd

type GetOutputFilename func(inputFilePath string) string
