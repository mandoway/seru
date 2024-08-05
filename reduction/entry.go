package reduction

import (
	"fmt"
	"github.com/mandoway/seru/files"
	"log"
	"os"
	"path"
	"time"
)

func StartReductionProcess(inputFile, testScript string) error {
	log.Println("Starting Reduction Process")
	log.Println("InputFile:", inputFile)
	log.Println("TestScript:", testScript)

	tmpDir := fmt.Sprintf("seru_reduction_%s", time.Now().Format(time.RFC3339))
	err := os.Mkdir(tmpDir, 0750)
	if err != nil {
		return err
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tmpDir)

	_, inputFileName := path.Split(inputFile)
	_, testScriptName := path.Split(testScript)

	err = files.Copy(path.Join(tmpDir, inputFileName), inputFile)
	if err != nil {
		return err
	}
	err = files.Copy(path.Join(tmpDir, testScriptName), testScript)
	if err != nil {
		return err
	}

	return nil
}
