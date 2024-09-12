package candidate

import "slices"

func MinCandidate(candidates []CandidateWithSize) CandidateWithSize {
	return slices.MinFunc(candidates, func(a, b CandidateWithSize) int {
		return b.Size - a.Size
	})
}
