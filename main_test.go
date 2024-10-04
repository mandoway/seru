package main

import (
	"github.com/mandoway/seru/reduction"
	"github.com/mandoway/seru/reduction/context"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestReduction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	aCtx := context.NewAlgorithmContext(false)
	dir := path.Join("test", "instances", "error", "issue_2246", "v1")
	ctx, err := context.NewRunContext("cue", path.Join(dir, "in.cue"), path.Join(dir, "test.sh"), *aCtx)
	if err != nil {
		t.Fatal(err)
	}

	err = reduction.RunMainReductionLoop(ctx)
	if err != nil {
		t.Fatal(err)
	}

	reducedFile := ctx.BestResult().InputPath
	if _, err := os.Stat(reducedFile); err == nil {
		matches, err := filepath.Glob(context.RunContextFolderPrefix + "*")
		t.Logf("Files successfully created, deleting %d matches...", len(matches))
		if err != nil {
			return
		}
		for _, match := range matches {
			t.Logf("Deleting %s", match)
			err := os.RemoveAll(match)
			if err != nil {
				return
			}
		}
	} else {
		t.Errorf("File does not exist, %s", reducedFile)
	}
}
