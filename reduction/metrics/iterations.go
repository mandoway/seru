package metrics

import "time"

type Iterations struct {
	Items []*Iteration
}

func NewIterations() *Iterations {
	return &Iterations{Items: []*Iteration{}}
}

func (i *Iterations) AddIteration() {
	i.Items = append(i.Items, NewIteration())
}

func (i *Iterations) Current() *Iteration {
	return i.Items[len(i.Items)-1]
}

type Iteration struct {
	Counts                          *Counts
	BeforeSize, AfterSize           int
	SyntacticMillis, SemanticMillis int64
}

func NewIteration() *Iteration {
	return &Iteration{
		Counts: NewCounts(),
	}
}

func (i *Iteration) TrackSyntactic(start time.Time) {
	elapsed := time.Since(start)
	i.SyntacticMillis = elapsed.Milliseconds()
}

func (i *Iteration) TrackSemantic(start time.Time) {
	elapsed := time.Since(start)
	i.SemanticMillis = elapsed.Milliseconds()
}
