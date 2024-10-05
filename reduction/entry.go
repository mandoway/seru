package reduction

import (
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/syntactic"
)

func StartReductionProcess(inputFile, testScript, givenLanguage string, isolation bool, reducer syntactic.Reducer) error {
	logging.Default.Println("SeRu - Syntactic & Semantic Reduction")
	logging.Default.Println()
	logging.Default.Printf("Creating new run context with (input=%s, test=%s, lang=%s)\n", inputFile, testScript, givenLanguage)
	algorithmConfig := context.NewAlgorithmConfig(isolation, reducer)
	logging.Default.Printf("Running algorithm with config %s\n", algorithmConfig)

	runCtx, err := context.NewRunContext(givenLanguage, inputFile, testScript, *algorithmConfig)
	if err != nil {
		return err
	}

	err = RunMainReductionLoop(runCtx)
	if err != nil {
		return err
	}

	return nil
}
