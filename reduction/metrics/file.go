package metrics

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

func StoreMetrics(reductionDir, inputDir string, iterations *Iterations, strategyNames []string, totalDuration time.Duration) error {
	aggregated := NewJsonFormat(iterations, totalDuration, strategyNames, inputDir)
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
}

func NewJsonFormat(iterations *Iterations, totalDuration time.Duration, strategyNames []string, inputDir string) *JsonFormat {
	total := NewIteration()
	total.BeforeSize = iterations.Items[0].BeforeSize
	total.AfterSize = iterations.Current().AfterSize

	for _, item := range iterations.Items {
		total.SemanticMillis += item.SemanticMillis
		total.SyntacticMillis += item.SyntacticMillis
		for _, strategy := range strategyNames {
			total.Counts.TotalCandidatesByStrategy[strategy] += item.Counts.TotalCandidatesByStrategy[strategy]
			total.Counts.ValidCandidatesByStrategy[strategy] += item.Counts.ValidCandidatesByStrategy[strategy]
			total.Counts.TestScriptExecutionsByStrategy[strategy] += item.Counts.TestScriptExecutionsByStrategy[strategy]
		}
	}

	return &JsonFormat{
		Iterations:      iterations,
		Total:           total,
		TotalTimeMillis: totalDuration.Milliseconds(),
		TotalIterations: len(iterations.Items),
		InputDir:        inputDir,
	}
}
