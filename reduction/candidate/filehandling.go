package candidate

import (
	"errors"
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/domain"
	"github.com/mandoway/seru/reduction/logging"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var candidateDirectoryPrefix = "strategy"

func CheckAndKeepValidCandidates(candidates [][]byte, ctx *context.RunContext) []*domain.CandidateWithSize {
	var validCandidates []*domain.CandidateWithSize
	for i, candidate := range candidates {
		candidateFiles, err := writeCandidate(ctx, i, candidate)
		if err != nil {
			logging.LogSemantic("Error writing candidate, try next", err)
			continue
		}

		ok, err := testCandidate(candidateFiles)
		if err != nil {
			logging.LogSemantic("Error testing candidate", err)
		}
		if ok {
			size := ctx.CountTokens(candidate)
			validCandidates = append(validCandidates, candidateFiles.WithSize(size))
		}
	}

	return validCandidates
}

func DeleteAllCandidates() {
	matches, _ := filepath.Glob(candidateDirectoryPrefix + "*")
	if len(matches) > 0 {
		for _, match := range matches {
			_ = os.RemoveAll(match)
		}
	}
}

func writeCandidate(ctx *context.RunContext, i int, candidate []byte) (*domain.Candidate, error) {
	folderName := fmt.Sprintf("%s%d_candidate%d", candidateDirectoryPrefix, ctx.CurrentSemanticStrategy, i)
	dir := path.Join(ctx.ReductionDir, folderName)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return nil, err
	}
	newInputFile := path.Join(dir, ctx.InputFilename())
	newTestFile := path.Join(dir, ctx.TestFilename())

	err = os.WriteFile(newInputFile, candidate, 0750)
	if err != nil {
		return nil, err
	}

	err = files.Copy(ctx.Current.TestPath, newTestFile)
	if err != nil {
		return nil, err
	}
	return domain.NewCandidate(newInputFile, newTestFile), nil
}

func testCandidate(candidate *domain.Candidate) (bool, error) {
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

func buildTestCommand(candidate *domain.Candidate) exec.Cmd {
	dir, file := path.Split(candidate.TestPath)
	cmd := *exec.Command(fmt.Sprintf("./%s", file))
	cmd.Dir = dir
	return cmd
}
