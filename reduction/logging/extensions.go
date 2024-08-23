package logging

import (
	"github.com/mandoway/seru/reduction/context"
	"log"
)

func LogStartReduction(dir string, startSize int) {
	log.Println("Starting reduction loop")
	log.Println("Results will be created in", dir)
	log.Println("Initial token size of program:", startSize)

	// ... print further configuration details if there are any
}

func LogEndReduction(sizes context.SizeContext) {
	log.Println("Finished reduction loop")
	log.Printf("Reduced program to %d/%d tokens (%s)\n", sizes.BestSizeInTokens, sizes.StartSizeInTokens, sizes.AsPercent())
}

func LogSyntactic(msg ...any) {
	logWithPrefix("SYNTACTIC", msg...)
}

func LogSemantic(msg ...any) {
	logWithPrefix("SEMANTIC", msg...)
}

func logWithPrefix(prefix any, msg ...any) {
	args := append([]any{prefix}, msg...)
	log.Println(args...)
}
