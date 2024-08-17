package context

import (
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/semantic"
	"github.com/mandoway/seru/reduction/semantic/semantic_plugin"
	"github.com/mandoway/seru/reduction/syntactic/syntactic_reducers"
	"os"
	"path"
	"time"
)

const ReductionFolderPrefix = "seru_reduction_"

type RunContext struct {
	Language         string
	InputFile        string
	TestScript       string
	ReductionDir     string
	Sizes            SizeContext
	SyntacticReducer syntactic_reducers.Functions
	SemanticReducer  semantic.Functions
}

func (ctx RunContext) InputFilePath() string {
	return path.Join(ctx.ReductionDir, ctx.InputFile)
}

func (ctx RunContext) TestScriptPath() string {
	return path.Join(ctx.ReductionDir, ctx.TestScript)
}

func NewRunContext(givenLanguage, inputFile, testScript string) (*RunContext, error) {
	reductionDir := fmt.Sprintf("%s%s", ReductionFolderPrefix, time.Now().Format(time.RFC3339))
	err := os.Mkdir(reductionDir, 0750)
	if err != nil {
		return nil, err
	}

	inputFileName := path.Base(inputFile)
	testScriptName := path.Base(testScript)

	err = files.Copy(path.Join(reductionDir, inputFileName), inputFile)
	if err != nil {
		return nil, err
	}
	err = files.Copy(path.Join(reductionDir, testScriptName), testScript)
	if err != nil {
		return nil, err
	}

	language := takeLanguageOrDefault(givenLanguage, inputFile)

	semanticFunctions, err := semantic_plugin.LoadSemanticReductionPlugin(language)
	if err != nil {
		return nil, &RunContextErr{message: err.Error()}
	}

	currentSize, err := getStartSizeInTokens(inputFile, semanticFunctions.CountFunction)
	if err != nil {
		return nil, &RunContextErr{message: err.Error()}
	}
	sizeContext := SizeContext{
		StartSizeInTokens: currentSize,
		BestSizeInTokens:  currentSize,
	}

	// TODO determine syntactic reducer from config
	syntacticFunctions := syntactic_reducers.PersesReducerFunctions

	return &RunContext{
		Language:         language,
		InputFile:        inputFileName,
		TestScript:       testScriptName,
		ReductionDir:     reductionDir,
		Sizes:            sizeContext,
		SyntacticReducer: syntacticFunctions,
		SemanticReducer:  *semanticFunctions,
	}, nil
}

func takeLanguageOrDefault(language, file string) string {
	if language != "" {
		return language
	}

	return path.Ext(file)[1:]
}

func getStartSizeInTokens(inputFile string, countFunction semantic.CountFunctionType) (int, error) {
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}

	return countFunction(bytes), nil
}

type SizeContext struct {
	StartSizeInTokens int
	BestSizeInTokens  int
}

type RunContextErr struct {
	message string
}

func (e *RunContextErr) Error() string {
	return fmt.Sprintf("Error during creation of run context: %s", e.message)
}
