package candidate

import (
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
			logging.LogSemantic("Error writing candidate, try next")
			continue
		}

		ok := testCandidate(candidateFiles)
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
	newInputFile := path.Join(ctx.ReductionDir, folderName, ctx.InputFilename())
	newTestFile := path.Join(ctx.ReductionDir, folderName, ctx.TestFilename())

	err := os.WriteFile(newInputFile, candidate, 0750)
	if err != nil {
		return nil, err
	}

	err = files.Copy(ctx.Current.TestPath, newTestFile)
	if err != nil {
		return nil, err
	}
	return domain.NewCandidate(newInputFile, newTestFile), nil
}

func testCandidate(candidate *domain.Candidate) bool {
	cmd := buildTestCommand(candidate)
	err := cmd.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}

func buildTestCommand(candidate *domain.Candidate) exec.Cmd {
	return *exec.Command(candidate.TestPath)
}
