package reduction

import (
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/context"
	"github.com/mandoway/seru/reduction/syntactic"
	"log"
	"os"
	"path"
	"time"
)

func StartReductionProcess(inputFile, testScript, givenLanguage string) error {
	log.Println("Starting Reduction Process")
	log.Println("InputFile:", inputFile)
	log.Println("TestScript:", testScript)

	runCtx, err := setupReductionRun(inputFile, testScript, givenLanguage)
	if err != nil {
		return err
	}

	resultFile, err := syntactic.ReduceSyntactically(runCtx)
	if err != nil {
		return err
	}

	log.Println("ResultFile:", resultFile)

	return nil
}

func setupReductionRun(inputFile, testScript, givenLanguage string) (*context.RunContext, error) {
	tmpDir := fmt.Sprintf("seru_reduction_%s", time.Now().Format(time.RFC3339))
	err := os.Mkdir(tmpDir, 0750)
	if err != nil {
		return nil, err
	}

	inputFileName := path.Base(inputFile)
	testScriptName := path.Base(testScript)

	err = files.Copy(path.Join(tmpDir, inputFileName), inputFile)
	if err != nil {
		return nil, err
	}
	err = files.Copy(path.Join(tmpDir, testScriptName), testScript)
	if err != nil {
		return nil, err
	}

	language := takeLanguageOrDefault(givenLanguage, inputFile)

	return &context.RunContext{
		Language:     language,
		InputFile:    inputFileName,
		TestScript:   testScriptName,
		ReductionDir: tmpDir,
		BestSize:     0, // todo
	}, nil
}

func takeLanguageOrDefault(language, file string) string {
	if language != "" {
		return language
	}

	return path.Ext(file)[1:]
}
