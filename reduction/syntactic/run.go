package syntactic

import (
	"errors"
	"fmt"
	"github.com/mandoway/seru/reduction"
	"os"
	"path"
)

func ReduceSyntactically(ctx reduction.RunContext) (string, error) {
	// Todo generify for other reducers
	cmd := BuildPersesReductionCommand(ctx.InputFile, ctx.TestScript, ctx.Language)

	// Todo add time measurement

	fmt.Println("Running command:", cmd)

	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New("syntactic reduction failed: " + err.Error())
	}

	dir, file := path.Split(ctx.InputFile)
	reducedFilePath := path.Join(dir, "perses_result", file)

	return reducedFilePath, nil
}
