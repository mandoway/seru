package syntactic

import (
	"errors"
	"github.com/mandoway/seru/reduction/context"
	"log"
	"os"
	"path"
)

func ReduceSyntactically(ctx *context.RunContext) (string, error) {
	// Todo generify for other reducers
	cmd := BuildPersesReductionCommand(ctx.InputFilePath(), ctx.TestScriptPath(), ctx.Language)

	// Todo add time measurement

	log.Println("Running command:", cmd)

	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New("syntactic reduction failed: " + err.Error())
	}

	reducedFilePath := path.Join(ctx.ReductionDir, "perses_result", ctx.InputFile)

	return reducedFilePath, nil
}
