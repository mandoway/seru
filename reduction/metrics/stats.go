package metrics

import "time"

type StatsByStrategy struct {
	TestScriptExecutionsByStrategy map[string]int
	TestScriptMillisByStrategy     map[string]int64
	GenerationMicrosByStrategy     map[string]int64

	TotalCandidatesByStrategy   map[string]int
	ValidCandidatesByStrategy   map[string]int
	AppliedCandidatesByStrategy map[string]int
}

func NewStatsByStrategy() *StatsByStrategy {
	return &StatsByStrategy{
		TestScriptExecutionsByStrategy: make(map[string]int),
		TestScriptMillisByStrategy:     make(map[string]int64),
		GenerationMicrosByStrategy:     make(map[string]int64),

		TotalCandidatesByStrategy:   make(map[string]int),
		ValidCandidatesByStrategy:   make(map[string]int),
		AppliedCandidatesByStrategy: make(map[string]int),
	}
}

func (c *StatsByStrategy) IncrementTestScriptExecutionByStrategy(strategy string) {
	c.TestScriptExecutionsByStrategy[strategy]++
}

func (c *StatsByStrategy) IncrementTestScriptTimeByStrategy(strategyName string, duration time.Duration) {
	c.TestScriptMillisByStrategy[strategyName] += duration.Milliseconds()
}

func (c *StatsByStrategy) IncrementGenerationTimeByStrategy(strategyName string, duration time.Duration) {
	c.GenerationMicrosByStrategy[strategyName] += duration.Microseconds()
}

func (c *StatsByStrategy) AddCandidatesByStrategy(strategy string, total, valid int) {
	c.TotalCandidatesByStrategy[strategy] += total
	c.ValidCandidatesByStrategy[strategy] += valid
}

func (c *StatsByStrategy) IncrementAppliedCandidatesByStrategy(strategy string) {
	c.AppliedCandidatesByStrategy[strategy]++
}
