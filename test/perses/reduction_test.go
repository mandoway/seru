package perses

import (
	"github.com/mandoway/seru/reduction"
	"os"
	"path/filepath"
	"testing"
)

func TestReduction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	ctx, err := reduction.NewRunContext("cue", "issue2246_v1/in.cue", "issue2246_v1/test.sh")
	if err != nil {
		t.Fatal(err)
	}

	err = reduction.RunMainReductionLoop(ctx)
	if err != nil {
		t.Fatal(err)
	}

	reducedFile := ctx.SyntacticReducer.GetOutputFilename(ctx.Current.InputPath)
	if _, err := os.Stat(reducedFile); err == nil {
		matches, err := filepath.Glob(reduction.RunContextFolderPrefix + "*")
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
