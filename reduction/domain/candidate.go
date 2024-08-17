package domain

type Candidate struct {
	InputPath string
	TestPath  string
}

func NewCandidate(inputPath string, testPath string) *Candidate {
	return &Candidate{InputPath: inputPath, TestPath: testPath}
}
