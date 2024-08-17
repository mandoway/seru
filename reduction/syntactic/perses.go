package syntactic

import (
	"fmt"
	"github.com/mandoway/seru/reduction/domain"
	"os/exec"
	"path"
	"strings"
)

var PersesReducerFunctions = Functions{
	BuildReductionCommand: buildPersesReductionCommand,
	GetOutputFilename:     buildPersesOutputFilename,
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

func buildPersesOutputFilename(inputFilePath string) string {
	dir, file := path.Split(inputFilePath)
	return path.Join(dir, "perses_result", file)
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
