package syntactic

import (
	"github.com/mandoway/seru/files"
	"github.com/mandoway/seru/reduction/candidate"
	"os/exec"
)

var VulcanReducerFunctions = Functions{
	BuildReductionCommand: buildVulcanReductionCommand,
	GetOutputFilename: func(inputFilePath string) string {
		return files.AddFolderToFilePath(inputFilePath, "perses_result")
	},
	Init: initialisePerses,
	name: "Vulcan",
}

func buildVulcanReductionCommand(candidate candidate.Candidate, language string) *exec.Cmd {
	args := []string{
		"-jar", persesJar,
		"-i", candidate.InputPath,
		"-t", candidate.TestPath,
		"--enable-vulcan", "true",
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
