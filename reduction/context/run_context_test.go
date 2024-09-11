package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTakeLanguageGiven(t *testing.T) {
	result := takeLanguageOrDefaultToFileExt("cue", "file.c")
	assert.Equal(t, "cue", result)
}

func TestTakeLanguageFromFile(t *testing.T) {
	result := takeLanguageOrDefaultToFileExt("", "file.c")
	assert.Equal(t, "c", result)
}
