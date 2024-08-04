package syntactic

import (
	"fmt"
	"os/exec"
	"strings"
)

func BuildPersesReductionCommand(inputFile, testFile, language string) *exec.Cmd {
	args := []string{
		"-jar", "perses_deploy.jar",
		"-i", inputFile,
		"-t", testFile,
		"--call-formatter", "true",
	}

	if language != "" {
		args = append(args, "--language-ext-jars", toLanguageJar(language))
	}

	return exec.Command(
		"java",
		args...,
	)
}

func toLanguageJar(language string) string {
	return fmt.Sprintf("%s.jar", strings.ToLower(language))
}
