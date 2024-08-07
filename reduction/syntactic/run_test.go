package syntactic

import (
	"github.com/mandoway/seru/reduction/context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildPersesReductionCommand(t *testing.T) {
	ctx := context.RunContext{
		Language:   "cue",
		InputFile:  "in.cue",
		TestScript: "test.sh",
	}

	command := BuildPersesReductionCommand(ctx.InputFile, ctx.TestScript, ctx.Language)

	assert.ElementsMatch(t, command.Args, []string{"java", "-jar", "perses_deploy.jar", "-i", "in.cue", "-t", "test.sh", "--call-formatter", "true", "--language-ext-jars", "cue.jar"})
}

func TestBuildPersesReductionCommandWithoutLanguage(t *testing.T) {
	ctx := context.RunContext{
		Language:   "",
		InputFile:  "in.cue",
		TestScript: "test.sh",
	}

	command := BuildPersesReductionCommand(ctx.InputFile, ctx.TestScript, ctx.Language)

	assert.ElementsMatch(t, command.Args, []string{"java", "-jar", "perses_deploy.jar", "-i", "in.cue", "-t", "test.sh", "--call-formatter", "true"})
}
