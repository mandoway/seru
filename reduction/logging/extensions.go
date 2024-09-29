package logging

import (
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/reduction/context"
)

func LogStartReduction(dir string, startSize int) {
	Default.Println("Starting reduction loop")
	Default.Println("Results will be created in", dir)
	Default.Println("Initial token size of program:", startSize)

	// ... print further configuration details if there are any
}

func LogEndReduction(sizes context.SizeContext, result *candidate.CandidateWithSize) {
	Default.Println("Finished reduction loop")
	Default.Printf("Reduced program to %d/%d tokens (%s)\n", sizes.BestSizeInTokens, sizes.StartSizeInTokens, sizes.AsPercent())
	Default.Printf("Final result is located at %s", result.InputPath)
}
