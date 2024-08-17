package metrics

import (
	"github.com/mandoway/seru/reduction/semantic"
	"os"
)

func GetTokenSizeOfFile(inputPath string, counter semantic.TokenCountFunctionType) (int, error) {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		return 0, err
	}

	return counter(bytes), nil
}
