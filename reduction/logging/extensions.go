package logging

import (
	"fmt"
	"github.com/mandoway/seru/reduction/candidate"
)

func LogStartReduction(dir string, startSize int) {
	Default.Println("Starting reduction loop")
	Default.Println("Results will be created in", dir)
	Default.Println("Initial token size of program:", startSize)

	// ... print further configuration details if there are any
}

func LogEndReduction(startSize int, result *candidate.CandidateWithSize) {
	Default.Println("Finished reduction loop")
	relativeReduction := fmt.Sprintf("%.2f%%", float32(result.Size)/float32(startSize)*100)
	Default.Printf("Reduced program to %d/%d tokens (%s)\n", result.Size, startSize, relativeReduction)
	Default.Printf("Final result is located at %s", result.InputPath)
}
