package persistance

import (
	"fmt"
	"github.com/mandoway/seru/reduction/candidate"
	"path"
	"testing"
)

func TestExitCode(t *testing.T) {
	instances := map[string]bool{
		inTestAssetsPath("exitCode0.sh"): true,
		inTestAssetsPath("exitCode1.sh"): false,
	}

	for testScriptPath, expected := range instances {
		t.Run(fmt.Sprint(expected), func(t *testing.T) {
			actual, _ := testCandidate(candidate.NewCandidate("", testScriptPath))

			if actual != expected {
				t.Errorf("Expected %v, got %v", expected, actual)
			}
		})
	}
}

func inTestAssetsPath(file string) string {
	return path.Join("..", "..", "..", "test", "assets", file)
}
