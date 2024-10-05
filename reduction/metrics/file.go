package metrics

import (
	"encoding/json"
	"os"
	"path"
)

func StoreMetrics(dir string, iterations *Iterations) error {
	raw, err := json.Marshal(iterations)
	if err != nil {
		return err
	}

	file := path.Join(dir, "metrics.json")
	err = os.WriteFile(file, raw, 0666)

	return err
}
