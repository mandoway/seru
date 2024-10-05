package reduction

import (
	"bytes"
	"errors"
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/syntactic"
)

func ReduceSyntactically(candidate candidate.Candidate, reducerConfig syntactic.Functions, language string) (string, error) {
	// Todo add time measurement
	// todo extract metrics from perses

	reductionCmd := reducerConfig.BuildReductionCommand(candidate, language)

	logging.Syntactic.Println("Executing command:", reductionCmd)

	var stdout, stderr bytes.Buffer
	reductionCmd.Stderr = &stderr
	reductionCmd.Stdout = &stdout
	err := reductionCmd.Run()
	if err != nil {
		logging.Syntactic.Println(stderr.String())
		return "", errors.New("syntactic reduction failed: " + err.Error())
	}

	outputFilename := reducerConfig.GetOutputFilename(candidate.InputPath)

	return outputFilename, nil
}
