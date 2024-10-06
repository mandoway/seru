package syntactic

import (
	"fmt"
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/candidate"
	"github.com/mandoway/seru/tools"
	"os/exec"
	"strings"
)

const persesJar = "perses_deploy.jar"

var PersesReducerFunctions = Functions{
	BuildReductionCommand: buildPersesReductionCommand,
	GetOutputFilename: func(inputFilePath string) string {
		return files.AddFolderToFilePath(inputFilePath, "perses_result")
	},
	Init: initialisePerses,
	name: "Perses",
}

func buildPersesReductionCommand(candidate candidate.Candidate, language string) *exec.Cmd {
	args := []string{
		"-jar", persesJar,
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

func initialisePerses(language string) error {
	// Check if java is installed
	_, err := exec.LookPath("java")
	if err != nil {
		return err
	}

	err = tools.EnsureFileExists(persesJar)
	if err != nil {
		return err
	}

	languageSpecificJar := toLanguageJar(language)
	err = tools.EnsureFileExists(languageSpecificJar)
	if err != nil {
		return err
	}

	return nil
}
