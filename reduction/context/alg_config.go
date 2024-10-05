package context

import (
	"fmt"
	"github.com/mandoway/seru/reduction/syntactic"
)

type SemanticApplicationMethod string

const (
	ApplyFirstOnly   SemanticApplicationMethod = "ApplyFirstOnly"
	ApplyAllCombined SemanticApplicationMethod = "ApplyAllCombined"
)

type AlgorithmConfig struct {
	applicationMethod    SemanticApplicationMethod
	syntacticReducer     syntactic.Functions
	syntacticReducerName syntactic.Reducer
}

func (a AlgorithmConfig) String() string {
	return fmt.Sprintf("SemanticApplicationMethod: %s, SyntacticReducer: %s", a.applicationMethod, a.syntacticReducerName)
}

func NewAlgorithmConfig(useIsolation bool, reducer syntactic.Reducer) *AlgorithmConfig {
	var application SemanticApplicationMethod
	if useIsolation {
		application = ApplyFirstOnly
	} else {
		application = ApplyAllCombined
	}

	var syntacticReducer syntactic.Functions
	switch reducer {
	case syntactic.Perses:
		syntacticReducer = syntactic.PersesReducerFunctions
		break
	case syntactic.Vulcan:
		syntacticReducer = syntactic.VulcanReducerFunctions
		break
	}

	return &AlgorithmConfig{applicationMethod: application, syntacticReducer: syntacticReducer}
}
