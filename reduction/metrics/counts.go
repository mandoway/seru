package metrics

type Counts struct {
	TestScriptExecutionsByStrategy map[int]int
	TotalCandidatesByStrategy      map[int]int
	ValidCandidatesByStrategy      map[int]int
}

func NewCounts() *Counts {
	return &Counts{
		TestScriptExecutionsByStrategy: make(map[int]int),
		TotalCandidatesByStrategy:      make(map[int]int),
		ValidCandidatesByStrategy:      make(map[int]int),
	}
}

func (c *Counts) AddCandidatesByStrategy(strategyIndex, total, valid int) {
	c.TotalCandidatesByStrategy[strategyIndex] = total
	c.ValidCandidatesByStrategy[strategyIndex] = valid
}

func (c *Counts) IncrementTestScriptExecutionByStrategy(strategyIndex int) {
	if _, found := c.TestScriptExecutionsByStrategy[strategyIndex]; !found {
		c.TestScriptExecutionsByStrategy[strategyIndex] = 0
	}
	c.TestScriptExecutionsByStrategy[strategyIndex]++
}
