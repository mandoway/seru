package candidate

import (
	"github.com/mandoway/seru/util/collection"
	"slices"
)

func MinCandidate(candidates []CandidateWithSize) CandidateWithSize {
	return slices.MinFunc(candidates, func(a, b CandidateWithSize) int {
		return b.Size - a.Size
	})
}

func MinCandidateP(candidates []*CandidateWithSize) *CandidateWithSize {
	nonNilCandidates := collection.FilterSlice(candidates, func(it *CandidateWithSize) bool {
		return it != nil
	})
	return slices.MinFunc(nonNilCandidates, func(a, b *CandidateWithSize) int {
		return b.Size - a.Size
	})
}
