package reduction

import (
	"bytes"
	"errors"
	"github.com/mandoway/seru/reduction/domain"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/syntactic"
)

func ReduceSyntactically(candidate domain.Candidate, reducerConfig syntactic.Functions, language string) (string, error) {
	// Todo add time measurement
	// todo extract metrics from perses
	// todo print stdout if configured

	reductionCmd := reducerConfig.BuildReductionCommand(candidate, language)

	logging.LogSyntactic("Executing command:", reductionCmd)

	var stdout, stderr bytes.Buffer
	reductionCmd.Stderr = &stderr
	reductionCmd.Stdout = &stdout
	err := reductionCmd.Run()
	if err != nil {
		logging.LogSyntactic(stderr.String())
		return "", errors.New("syntactic reduction failed: " + err.Error())
	}

	outputFilename := reducerConfig.GetOutputFilename(candidate.InputPath)

	return outputFilename, nil
}
