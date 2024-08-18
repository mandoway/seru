package metrics

import (
	"github.com/mandoway/seru/reduction/plugin"
	"os"
)

func GetTokenSizeOfFile(inputPath string, counter plugin.TokenCountFunction) (int, error) {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		return 0, err
	}

	return counter(bytes), nil
}
