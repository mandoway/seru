package cmd

import (
	"github.com/mandoway/seru/reduction/logging"
	"github.com/mandoway/seru/reduction/syntactic"
)

type reductionType = string

const (
	PersesReducer reductionType = "perses"
	VulcanReducer reductionType = "vulcan"
)

type Flags struct {
	InputFile, TestScript, GivenLanguage string
	UseStrategyIsolation                 bool
	SyntacticReducer                     reductionType
}

func (f Flags) GetReducer() syntactic.Reducer {
	switch f.SyntacticReducer {
	case PersesReducer:
		return syntactic.Perses
	case VulcanReducer:
		return syntactic.Vulcan
	default:
		logging.Default.Printf("Unknown reducer %s, using Perses instead (%s|%s)", f.SyntacticReducer, PersesReducer, VulcanReducer)
		return syntactic.Perses
	}
}
