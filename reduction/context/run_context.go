package context

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/reduction/logging"
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

// RunContext holds all data necessary for reduction
// TODO add metrics like test script count
type RunContext struct {
	currentVersion int

	language                string
	reductionDir            string
	sizes                   SizeContext
	currentSemanticStrategy int
	semanticStrategiesTotal int

	algorithmContext                 AlgorithmContext
	forceExhaustedSemanticStrategies bool

	bestResult *candidate.CandidateWithSize
	hashOfBest [16]byte

	countTokens     plugin.TokenCountFunction
	getStrategyName plugin.StrategyNameFunction

	syntacticReducer syntactic.Functions

	semanticReducer plugin.SemanticReductionFunction
}

func (ctx *RunContext) BestResult() *candidate.CandidateWithSize {
	return ctx.bestResult
}

func (ctx *RunContext) Sizes() SizeContext {
	return ctx.sizes
}

func (ctx *RunContext) CurrentSemanticStrategy() int {
	return ctx.currentSemanticStrategy
}

func (ctx *RunContext) IncrementSemanticStrategy() {
	ctx.currentSemanticStrategy++
}

func (ctx *RunContext) SemanticStrategiesTotal() int {
	return ctx.semanticStrategiesTotal
}

func (ctx *RunContext) GetStrategyName(index int) string {
	return ctx.getStrategyName(index)
}

func (ctx *RunContext) SemanticApplicationMethod() SemanticApplicationMethod {
	return ctx.algorithmContext.applicationMethod
}

func (ctx *RunContext) SetExhaustedSemanticStrategies() {
	if ctx.algorithmContext.applicationMethod == ApplyFirstOnly {
		panic("Forcing strategy exhaustion is only supported with combined strategies")
	}
	ctx.forceExhaustedSemanticStrategies = true
}

func (ctx *RunContext) ExhaustedSemanticStrategies() bool {
	switch ctx.algorithmContext.applicationMethod {
	case ApplyFirstOnly:
		return ctx.currentSemanticStrategy >= ctx.semanticStrategiesTotal
	case ApplyAllCombined:
		return ctx.forceExhaustedSemanticStrategies
	}

	panic(fmt.Sprintf("unknown algorithm method: %s", ctx.algorithmContext.applicationMethod))
}

func (ctx *RunContext) CountTokens(bytes []byte) int {
	return ctx.countTokens(bytes)
}

func (ctx *RunContext) SyntacticReducer() syntactic.Functions {
	return ctx.syntacticReducer
}

func (ctx *RunContext) SemanticReduce(bytes []byte) ([][]byte, error) {
	return ctx.semanticReducer(bytes, ctx.currentSemanticStrategy)
}

func (ctx *RunContext) SemanticReduceWithStrategy(bytes []byte, strategyIndex int) ([][]byte, error) {
	if strategyIndex >= ctx.semanticStrategiesTotal {
		return [][]byte{}, errors.New(fmt.Sprintf("no strategy available at index %d", strategyIndex))
	}

	return ctx.semanticReducer(bytes, strategyIndex)
}

func (ctx *RunContext) Language() string {
	return ctx.language
}

func (ctx *RunContext) ReductionDir() string {
	return ctx.reductionDir
}

func (ctx *RunContext) InputFilename() string {
	return path.Base(ctx.BestResult().InputPath)
}

func (ctx *RunContext) TestFilename() string {
	return path.Base(ctx.BestResult().TestPath)
}

func (ctx *RunContext) UpdateCurrent(candidatePath string, candidateSize int) error {
	logging.Default.Println("Store new best with size", candidateSize)
	err := files.Copy(candidatePath, ctx.BestResult().InputPath)
	if err != nil {
		return err
	}
	ctx.sizes.BestSizeInTokens = candidateSize
	ctx.bestResult.Size = candidateSize
	ctx.currentVersion++

	err = ctx.saveCurrent()
	if err != nil {
		return err
	}

	return nil
}

func (ctx *RunContext) saveCurrent() error {
	dir := path.Join(ctx.ReductionDir(), strconv.Itoa(ctx.currentVersion))
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}
	newInputFile := path.Join(dir, ctx.InputFilename())
	newTestFile := path.Join(dir, ctx.TestFilename())

	err = files.Copy(ctx.BestResult().InputPath, newInputFile)
	if err != nil {
		return err
	}

	err = files.Copy(ctx.BestResult().TestPath, newTestFile)
	if err != nil {
		return err
	}

	err = ctx.updateHash()
	return err
}

func (ctx *RunContext) updateHash() error {
	file, err := os.ReadFile(ctx.bestResult.InputPath)
	if err != nil {
		return err
	}
	ctx.hashOfBest = md5.Sum(file)

	return nil
}

func (ctx *RunContext) GetHash() [16]byte {
	return ctx.hashOfBest
}

func NewRunContext(givenLanguage, inputFilePath, testScriptPath string, algoContext AlgorithmContext) (*RunContext, error) {
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

	bestCandidate := candidate.NewCandidate(inputFileInReductionDir, testScriptInReductionDir).WithSize(currentSize)
	ctx := &RunContext{
		currentVersion: 0,

		language:                language,
		reductionDir:            reductionDir,
		sizes:                   sizeContext,
		semanticStrategiesTotal: semanticStrategiesSize,
		currentSemanticStrategy: 0,

		algorithmContext: algoContext,

		bestResult: bestCandidate,

		syntacticReducer: syntacticFunctions,
		semanticReducer:  pluginFunctions.SemanticReduce,
		countTokens:      pluginFunctions.CountTokens,
		getStrategyName:  pluginFunctions.GetStrategyName,
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
