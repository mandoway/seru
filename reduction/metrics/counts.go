package metrics

type Counts struct {
	TestScriptExecutionsByStrategy map[string]int
	TotalCandidatesByStrategy      map[string]int
	ValidCandidatesByStrategy      map[string]int
}

func NewCounts() *Counts {
	return &Counts{
		TestScriptExecutionsByStrategy: make(map[string]int),
		TotalCandidatesByStrategy:      make(map[string]int),
		ValidCandidatesByStrategy:      make(map[string]int),
	}
}

func (c *Counts) AddCandidatesByStrategy(strategy string, total, valid int) {
	c.TotalCandidatesByStrategy[strategy] += total
	c.ValidCandidatesByStrategy[strategy] += valid
}

func (c *Counts) IncrementTestScriptExecutionByStrategy(strategy string) {
	if _, found := c.TestScriptExecutionsByStrategy[strategy]; !found {
		c.TestScriptExecutionsByStrategy[strategy] = 0
	}
	c.TestScriptExecutionsByStrategy[strategy]++
}
