package reduction

import (
	"bytes"
	"errors"
	"github.com/mandoway/seru/reduction/domain"
	"github.com/mandoway/seru/reduction/syntactic"
	"log"
)

func ReduceSyntactically(candidate domain.Candidate, reducerConfig syntactic.Functions, language string) (string, error) {
	// Todo add time measurement
	// todo extract metrics from perses
	// todo print stdout if configured

	reductionCmd := reducerConfig.BuildReductionCommand(candidate, language)

	log.Println("Syntax reduction command:", reductionCmd)

	var stdout, stderr bytes.Buffer
	reductionCmd.Stderr = &stderr
	reductionCmd.Stdout = &stdout
	err := reductionCmd.Run()
	if err != nil {
		log.Println(stderr.String())
		return "", errors.New("syntactic reduction failed: " + err.Error())
	}

	outputFilename := reducerConfig.GetOutputFilename(candidate.InputPath)

	return outputFilename, nil
}
