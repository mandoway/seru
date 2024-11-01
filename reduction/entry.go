package reduction

import (
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/syntactic"
	"github.com/mandoway/seru/util/collection"
)

func StartReductionProcess(inputFile, testScript, givenLanguage string, isolation, enableMetrics bool, reducer syntactic.Reducer, activeStrategies *collection.Set) error {
	logging.Default.Println("SeRu - Syntactic & Semantic Reduction")
	logging.Default.Println()
	logging.Default.Printf("Creating new run context with (input=%s, test=%s, lang=%s)\n", inputFile, testScript, givenLanguage)
	algorithmConfig := context.NewAlgorithmConfig(isolation, reducer, activeStrategies)
	logging.Default.Printf("Running algorithm with config %s\n", algorithmConfig)

	runCtx, err := context.NewRunContext(givenLanguage, inputFile, testScript, *algorithmConfig)
	if err != nil {
		return err
	}

	err = RunMainReductionLoop(runCtx, enableMetrics)
	if err != nil {
		return err
	}

	return nil
}
