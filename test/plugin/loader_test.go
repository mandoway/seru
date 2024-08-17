package plugin

import (
	"github.com/mandoway/seru/reduction/semantic/semantic_plugin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPluginLoader(t *testing.T) {
	functions, err := semantic_plugin.LoadSemanticReductionPlugin("cue")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, functions)
	assert.NotNil(t, functions.ReduceFunction)
	assert.NotNil(t, functions.CountFunction)
}
