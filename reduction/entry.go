package reduction

import (
	"log"
)

func StartReductionProcess(inputFile, testScript, givenLanguage string) error {
	log.Println("SeRu - Syntactic & Semantic Reduction")
	log.Println()
	log.Printf("Creating new run context with (input=%s, test=%s, lang=%s)\n", inputFile, testScript, givenLanguage)

	runCtx, err := NewRunContext(givenLanguage, inputFile, testScript)
	if err != nil {
		return err
	}

	err = RunMainReductionLoop(runCtx)
	if err != nil {
		return err
	}

	return nil
}
