package reduction

import (
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/syntactic"
	"log"
)

func StartReductionProcess(inputFile, testScript, givenLanguage string) error {
	log.Println("Starting Reduction Process")
	log.Println("InputFile:", inputFile)
	log.Println("TestScript:", testScript)

	runCtx, err := context.NewRunContext(givenLanguage, inputFile, testScript)
	if err != nil {
		return err
	}

	cmd := runCtx.SyntacticReducer.BuildReductionCommand(runCtx.InputFilePath(), runCtx.TestScriptPath(), runCtx.Language)
	err = syntactic.ReduceSyntactically(cmd)
	if err != nil {
		return err
	}

	//log.Println("ResultFile:", resultFile)

	return nil
}
