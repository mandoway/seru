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
	"strconv"
	"time"
)

const RunContextFolderPrefix = "seru_reduction_"

// TODO add metrics like test script count
type RunContext struct {
	currentVersion int

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

func (ctx *RunContext) InputFilename() string {
	return path.Base(ctx.Current.InputPath)
}

func (ctx *RunContext) TestFilename() string {
	return path.Base(ctx.Current.TestPath)
}

func (ctx *RunContext) UpdateCurrent(candidatePath string, candidateSize int) error {
	log.Println("Store new best with size", candidateSize)
	err := files.Copy(candidatePath, ctx.Current.InputPath)
	if err != nil {
		return err
	}
	ctx.Sizes.BestSizeInTokens = candidateSize
	ctx.currentVersion++

	err = ctx.saveCurrent()
	if err != nil {
		return err
	}

	return nil
}

func (ctx *RunContext) saveCurrent() error {
	dir := path.Join(ctx.ReductionDir, strconv.Itoa(ctx.currentVersion))
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}
	newInputFile := path.Join(dir, ctx.InputFilename())
	newTestFile := path.Join(dir, ctx.TestFilename())

	err = files.Copy(ctx.Current.InputPath, newInputFile)
	if err != nil {
		return err
	}

	err = files.Copy(ctx.Current.TestPath, newTestFile)
	if err != nil {
		return err
	}

	return nil
}

func NewRunContext(givenLanguage, inputFilePath, testScriptPath string) (*RunContext, error) {
	// Copy input files
	reductionDir := fmt.Sprintf("%s%s", RunContextFolderPrefix, time.Now().Format(time.RFC3339))
	err := os.Mkdir(reductionDir, 0750)
	if err != nil {
		return nil, NewRunContextErr(err)
	}

	inputFileInReductionDir := getPathInFolder(reductionDir, inputFilePath)
	testScriptInReductionDir := getPathInFolder(reductionDir, testScriptPath)

	err = files.Copy(inputFilePath, inputFileInReductionDir)
	if err != nil {
		return nil, NewRunContextErr(err)
	}
	err = files.Copy(testScriptPath, testScriptInReductionDir)
	if err != nil {
		return nil, NewRunContextErr(err)
	}

	language := takeLanguageOrDefaultToFileExt(givenLanguage, inputFilePath)

	// Semantic plugin reducer config
	pluginFunctions, err := plugin.LoadSemanticReductionPlugin(language)
	if err != nil {
		return nil, NewRunContextErr(err)
	}

	semanticStrategiesSize := pluginFunctions.GetStrategyCount()

	// TODO determine syntactic reducer from config
	// Syntactic reducer config
	syntacticFunctions := syntactic.PersesReducerFunctions
	err = syntacticFunctions.Init(language)
	if err != nil {
		return nil, NewRunContextErr(err)
	}

	// Size
	currentSize, err := metrics.GetTokenSizeOfFile(inputFilePath, pluginFunctions.CountTokens)
	if err != nil {
		return nil, NewRunContextErr(err)
	}
	sizeContext := SizeContext{
		StartSizeInTokens: currentSize,
		BestSizeInTokens:  currentSize,
	}

	ctx := &RunContext{
		Language:                language,
		Current:                 domain.NewCandidate(inputFileInReductionDir, testScriptInReductionDir),
		currentVersion:          0,
		ReductionDir:            reductionDir,
		Sizes:                   sizeContext,
		SyntacticReducer:        syntacticFunctions,
		SemanticReducer:         pluginFunctions.SemanticReduce,
		CountTokens:             pluginFunctions.CountTokens,
		SemanticStrategiesTotal: semanticStrategiesSize,
		CurrentSemanticStrategy: 0,
	}

	err = ctx.saveCurrent()
	if err != nil {
		return nil, NewRunContextErr(err)
	}

	return ctx, nil
}

func getPathInFolder(folder, filePath string) string {
	return path.Join(folder, path.Base(filePath))
}

func takeLanguageOrDefaultToFileExt(language, file string) string {
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

func NewRunContextErr(err error) *RunContextErr {
	return &RunContextErr{message: err.Error()}
}

func (e *RunContextErr) Error() string {
	return fmt.Sprintf("Error during creation of run context: %s", e.message)
}
