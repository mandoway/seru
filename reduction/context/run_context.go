package context

import (
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/domain"
	"github.com/mandoway/seru/reduction/metrics"
	"github.com/mandoway/seru/reduction/plugin"
	"github.com/mandoway/seru/reduction/syntactic"
	"log"
	"os"
	"path"
	"time"
)

const RunContextFolderPrefix = "seru_reduction_"

// TODO add metrics like test script count
type RunContext struct {
	Language                string
	Current                 *domain.Candidate
	ReductionDir            string
	Sizes                   SizeContext
	SyntacticReducer        syntactic.Functions
	CountTokens             plugin.TokenCountFunction
	SemanticReducer         plugin.SemanticReductionFunction
	CurrentSemanticStrategy int
	SemanticStrategiesTotal int
}

func (ctx RunContext) InputFilename() string {
	return path.Base(ctx.Current.InputPath)
}

func (ctx RunContext) TestFilename() string {
	return path.Base(ctx.Current.TestPath)
}

func NewRunContext(givenLanguage, inputFilePath, testScriptPath string) (*RunContext, error) {
	reductionDir := fmt.Sprintf("%s%s", RunContextFolderPrefix, time.Now().Format(time.RFC3339))
	err := os.Mkdir(reductionDir, 0750)
	if err != nil {
		return nil, err
	}

	inputFileInReductionDir := getPathInFolder(reductionDir, inputFilePath)
	testScriptInReductionDir := getPathInFolder(reductionDir, testScriptPath)

	err = files.Copy(inputFilePath, inputFileInReductionDir)
	if err != nil {
		return nil, err
	}
	err = files.Copy(testScriptPath, testScriptInReductionDir)
	if err != nil {
		return nil, err
	}

	language := takeLanguageOrDefault(givenLanguage, inputFilePath)

	pluginFunctions, err := plugin.LoadSemanticReductionPlugin(language)
	if err != nil {
		return nil, &RunContextErr{message: err.Error()}
	}

	currentSize, err := metrics.GetTokenSizeOfFile(inputFilePath, pluginFunctions.CountTokens)
	if err != nil {
		return nil, &RunContextErr{message: err.Error()}
	}
	sizeContext := SizeContext{
		StartSizeInTokens: currentSize,
		BestSizeInTokens:  currentSize,
	}

	// TODO determine syntactic reducer from config
	syntacticFunctions := syntactic.PersesReducerFunctions

	semanticStrategiesSize := pluginFunctions.GetStrategyCount()

	return &RunContext{
		Language:                language,
		Current:                 domain.NewCandidate(inputFileInReductionDir, testScriptInReductionDir),
		ReductionDir:            reductionDir,
		Sizes:                   sizeContext,
		SyntacticReducer:        syntacticFunctions,
		SemanticReducer:         pluginFunctions.SemanticReduce,
		CountTokens:             pluginFunctions.CountTokens,
		SemanticStrategiesTotal: semanticStrategiesSize,
		CurrentSemanticStrategy: 0,
	}, nil
}

func getPathInFolder(folder, filePath string) string {
	return path.Join(folder, path.Base(filePath))
}

func takeLanguageOrDefault(language, file string) string {
	if language != "" {
		return language
	}

	fileEnding := path.Ext(file)[1:]
	log.Printf("No language configured, using language from file '%s'\n", fileEnding)
	return fileEnding
}

type SizeContext struct {
	StartSizeInTokens int
	BestSizeInTokens  int
}

func (s SizeContext) AsPercent() string {
	return fmt.Sprintf("%.2f%%", float32(s.BestSizeInTokens)/float32(s.StartSizeInTokens)*100)
}

type RunContextErr struct {
	message string
}

func (e *RunContextErr) Error() string {
	return fmt.Sprintf("Error during creation of run context: %s", e.message)
}
