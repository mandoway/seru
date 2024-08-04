package perses

import (
	"github.com/mandoway/seru/reduction"
	"github.com/mandoway/seru/reduction/syntactic"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestSyntacticReduction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	ctx := reduction.RunContext{
		Language:   "cue",
		InputFile:  "issue2246_v1/in.cue",
		TestScript: "issue2246_v1/test.sh",
	}

	reducedFile, err := syntactic.ReduceSyntactically(ctx)
	if err != nil {
		t.Errorf("Failed with %v", err)
		return
	}

	if _, err := os.Stat(reducedFile); err == nil {
		t.Log("Files successfully created, deleting...")
		dir, _ := path.Split(reducedFile)
		err := os.RemoveAll(dir)
		if err != nil {
			return
		}
		files, err := filepath.Glob("issue2246_v1/*.orig")
		if err != nil {
			return
		}
		for _, file := range files {
			err := os.Remove(file)
			if err != nil {
				return
			}
		}
	} else {
		t.Errorf("File does not exist, %s", reducedFile)
	}
}
