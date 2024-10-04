package candidate

type Candidate struct {
	InputPath string
	TestPath  string
}

func (c Candidate) WithSize(size int) *CandidateWithSize {
	return &CandidateWithSize{
		Size:      size,
		Candidate: c,
	}
}

func NewCandidate(inputPath string, testPath string) *Candidate {
	return &Candidate{InputPath: inputPath, TestPath: testPath}
}

type CandidateWithSize struct {
	Size int
	Candidate
}
