package persistance

import (
	"errors"
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/logging"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var candidateDirectoryPrefix = "strategy"

func CheckAndKeepValidCandidates(candidates [][]byte, ctx *context.RunContext, currentStrategy int) []*candidate.CandidateWithSize {
	var validCandidates []*candidate.CandidateWithSize
	for i, currentCandidate := range candidates {
		candidateFiles, err := writeCandidate(ctx, i, currentCandidate, currentStrategy)
		if err != nil {
			logging.LogSemantic("Error writing candidate, try next", err)
			continue
		}

		ok, err := testCandidate(candidateFiles)
		if err != nil {
			logging.LogSemantic("Error testing candidate", err)
		}
		if ok {
			size := ctx.CountTokens(currentCandidate)
			validCandidates = append(validCandidates, candidateFiles.WithSize(size))
		} else {
			deleteCandidate(candidateFiles)
		}
	}

	return validCandidates
}

func deleteCandidate(candidateFiles *candidate.Candidate) {
	dir := path.Dir(candidateFiles.InputPath)
	_ = os.RemoveAll(dir)
}

func DeleteAllCandidates(ctx *context.RunContext) {
	reductionDir := ctx.ReductionDir()
	matches, _ := filepath.Glob(reductionDir + "/" + candidateDirectoryPrefix + "*")
	if len(matches) > 0 {
		for _, match := range matches {
			_ = os.RemoveAll(match)
		}
	}

	outputFileRelativeToReductionDir := ctx.SyntacticReducer().GetOutputFilename(ctx.InputFilename())
	syntacticReducerLeftoverDir := path.Join(ctx.ReductionDir(), path.Dir(outputFileRelativeToReductionDir))
	_ = os.RemoveAll(syntacticReducerLeftoverDir)
}

func writeCandidate(ctx *context.RunContext, i int, currentCandidate []byte, currentStrategy int) (*candidate.Candidate, error) {
	folderName := fmt.Sprintf("%s%d_candidate%d", candidateDirectoryPrefix, currentStrategy, i)
	dir := path.Join(ctx.ReductionDir(), folderName)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return nil, err
	}
	newInputFile := path.Join(dir, ctx.InputFilename())
	newTestFile := path.Join(dir, ctx.TestFilename())

	err = os.WriteFile(newInputFile, currentCandidate, 0750)
	if err != nil {
		return nil, err
	}

	err = files.Copy(ctx.BestResult().TestPath, newTestFile)
	if err != nil {
		return nil, err
	}
	return candidate.NewCandidate(newInputFile, newTestFile), nil
}

func testCandidate(candidate *candidate.Candidate) (bool, error) {
	cmd := buildTestCommand(candidate)
	err := cmd.Run()
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func buildTestCommand(candidate *candidate.Candidate) exec.Cmd {
	dir, file := path.Split(candidate.TestPath)
	cmd := *exec.Command(fmt.Sprintf("./%s", file))
	cmd.Dir = dir
	return cmd
}
