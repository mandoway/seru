package plugin

import (
	"github.com/mandoway/seru/reduction/semantic/plugin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPluginLoader(t *testing.T) {
	functions, err := plugin.LoadSemanticReductionPlugin("cue")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, functions)
	assert.NotNil(t, functions.MainFunction)
	assert.NotNil(t, functions.CountFunction)
}
