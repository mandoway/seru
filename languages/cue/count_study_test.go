package main

import (
	"fmt"
	"github.com/mandoway/seru/languages/cue/language"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCountEveryStudyFolder(t *testing.T) {

	resultFolder := "groundtruth_vulcan"
	studyFolder := "../../study/" + resultFolder + "/instances"

	var results []string
	err := filepath.WalkDir(studyFolder, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() || filepath.Ext(path) != ".cue" {
			return nil
		}

		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		issuePath, _ := strings.CutPrefix(path, studyFolder)
		parts := strings.Split(issuePath, string(os.PathSeparator))
		severity := parts[1]
		issue := parts[2]
		version := parts[3]
		run, _ := strings.CutPrefix(parts[4], "run_")
		tokens := language.CountTokensUsingScanner(file)
		line := fmt.Sprintf("%s,%s,%s,%s,%d", severity, issue, version, run, tokens)
		results = append(results, line)
		return nil
	})
	if err != nil {
		t.Fatalf("failed to read .cue files: %v", err)
	}

	t.Logf("Counting results: %v", results)

	outputFile, err := os.Create(resultFolder + "_tokens.csv")
	if err != nil {
		t.Fatalf("failed to create output.csv file: %v", err)
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString("severity,issue,version,run,tokens\n")
	if err != nil {
		return
	}
	for _, line := range results {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			t.Fatalf("failed to write to output.csv: %v", err)
		}
	}
}
