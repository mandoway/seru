package syntactic

import (
	"github.com/mandoway/seru/reduction/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildPersesReductionCommand(t *testing.T) {
	command := PersesReducerFunctions.BuildReductionCommand(*domain.NewCandidate("in.cue", "test.sh"), "cue")

	assert.ElementsMatch(t, command.Args, []string{"java", "-jar", "perses_deploy.jar", "-i", "in.cue", "-t", "test.sh", "--call-formatter", "true", "--language-ext-jars", "cue.jar"})
}

func TestBuildPersesReductionCommandWithoutLanguage(t *testing.T) {
	command := PersesReducerFunctions.BuildReductionCommand(*domain.NewCandidate("in.cue", "test.sh"), "")

	assert.ElementsMatch(t, command.Args, []string{"java", "-jar", "perses_deploy.jar", "-i", "in.cue", "-t", "test.sh", "--call-formatter", "true"})
}
