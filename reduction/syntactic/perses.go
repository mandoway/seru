package syntactic

import (
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/domain"
	"os/exec"
	"strings"
)

var PersesReducerFunctions = Functions{
	BuildReductionCommand: buildPersesReductionCommand,
	GetOutputFilename: func(inputFilePath string) string {
		return files.AddFolderToFilePath(inputFilePath, "perses_result")
	},
}

func buildPersesReductionCommand(candidate domain.Candidate, language string) *exec.Cmd {
	args := []string{
		"-jar", "perses_deploy.jar",
		"-i", candidate.InputPath,
		"-t", candidate.TestPath,
		"--call-formatter", "true",
	}

	if _, isSupported := languagesSupportedByPerses[language]; !isSupported && language != "" {
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

var languagesSupportedByPerses = map[string]struct{}{
	"c":  {},
	"cc": {}, "cpp": {}, "cxx": {},
	"rs":    {},
	"scala": {}, "sc": {},
	"java":       {},
	"javascript": {}, "js": {},
	"py": {}, "py3": {},
	"glsl": {}, "comp": {}, "frag": {}, "vert": {},
	"go":     {},
	"php":    {},
	"rb":     {},
	"sqlite": {},
	"mysql":  {},
	"sol":    {},
	"v":      {}, "sv": {},
	"smt2": {},
}
