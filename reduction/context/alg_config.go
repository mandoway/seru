package context

import (
	"fmt"
	"github.com/mandoway/seru/reduction/syntactic"
	"github.com/mandoway/seru/util/collection"
)

type SemanticApplicationMethod string

const (
	ApplyFirstOnly   SemanticApplicationMethod = "ApplyFirstOnly"
	ApplyAllCombined SemanticApplicationMethod = "ApplyAllCombined"
)

type AlgorithmConfig struct {
	applicationMethod SemanticApplicationMethod
	syntacticReducer  syntactic.Functions
	activeStrategies  *collection.Set
}

func (a AlgorithmConfig) String() string {
	activeStrategiesStr := "all"
	if !a.activeStrategies.IsEmpty() {
		activeStrategiesStr = a.activeStrategies.String()
	}
	return fmt.Sprintf("SemanticApplicationMethod: %s, SyntacticReducer: %s, Strategies: %s", a.applicationMethod, a.syntacticReducer, activeStrategiesStr)
}

func NewAlgorithmConfig(useIsolation bool, reducer syntactic.Reducer, activeStrategies *collection.Set) *AlgorithmConfig {
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

	return &AlgorithmConfig{applicationMethod: application, syntacticReducer: syntacticReducer, activeStrategies: activeStrategies}
}
