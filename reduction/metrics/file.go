package metrics

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

func StoreMetrics(reductionDir, inputDir, startTime string, iterations *Iterations, strategyNames []string, totalDuration time.Duration) error {
	aggregated := NewJsonFormat(iterations, totalDuration, strategyNames, inputDir, startTime)
	raw, err := json.MarshalIndent(aggregated, "", "  ")
	if err != nil {
		return err
	}

	file := path.Join(reductionDir, "metrics.json")
	err = os.WriteFile(file, raw, 0666)

	return err
}

type JsonFormat struct {
	Iterations      *Iterations
	Total           *Iteration
	TotalTimeMillis int64
	TotalIterations int
	InputDir        string
	StartTime       string
}

func NewJsonFormat(iterations *Iterations, totalDuration time.Duration, strategyNames []string, inputDir, startTime string) *JsonFormat {
	total := NewIteration()
	total.BeforeSize = iterations.Items[0].BeforeSize
	total.AfterSize = iterations.Current().AfterSize

	for _, item := range iterations.Items {
		total.SemanticMillis += item.SemanticMillis
		total.SyntacticMillis += item.SyntacticMillis
		for _, strategy := range strategyNames {
			total.StatsByStrategy.TestScriptExecutionsByStrategy[strategy] += item.StatsByStrategy.TestScriptExecutionsByStrategy[strategy]
			total.StatsByStrategy.TestScriptMillisByStrategy[strategy] += item.StatsByStrategy.TestScriptMillisByStrategy[strategy]
			total.StatsByStrategy.GenerationMicrosByStrategy[strategy] += item.StatsByStrategy.GenerationMicrosByStrategy[strategy]

			total.StatsByStrategy.TotalCandidatesByStrategy[strategy] += item.StatsByStrategy.TotalCandidatesByStrategy[strategy]
			total.StatsByStrategy.ValidCandidatesByStrategy[strategy] += item.StatsByStrategy.ValidCandidatesByStrategy[strategy]
			total.StatsByStrategy.AppliedCandidatesByStrategy[strategy] += item.StatsByStrategy.AppliedCandidatesByStrategy[strategy]
		}
	}

	return &JsonFormat{
		Iterations:      iterations,
		Total:           total,
		TotalTimeMillis: totalDuration.Milliseconds(),
		TotalIterations: len(iterations.Items),
		InputDir:        inputDir,
		StartTime:       startTime,
	}
}
