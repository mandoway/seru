package reduction

import (
	"github.com/mandoway/seru/reduction/context"
	"log"
)

func StartReductionProcess(inputFile, testScript, givenLanguage string, isolation bool) error {
	log.Println("SeRu - Syntactic & Semantic Reduction")
	log.Println()
	log.Printf("Creating new run context with (input=%s, test=%s, lang=%s)\n", inputFile, testScript, givenLanguage)
	algorithmContext := context.NewAlgorithmContext(isolation)
	log.Printf("Running algorithm with context %v\n", algorithmContext)

	runCtx, err := context.NewRunContext(givenLanguage, inputFile, testScript, *algorithmContext)
	if err != nil {
		return err
	}

	err = RunMainReductionLoop(runCtx)
	if err != nil {
		return err
	}

	return nil
}
