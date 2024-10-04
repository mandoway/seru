package context

type SemanticApplicationMethod string

const (
	ApplyFirstOnly   SemanticApplicationMethod = "ApplyFirstOnly"
	ApplyAllCombined SemanticApplicationMethod = "ApplyAllCombined"
)

type AlgorithmContext struct {
	applicationMethod SemanticApplicationMethod
}

func NewAlgorithmContext(useIsolation bool) *AlgorithmContext {
	var application SemanticApplicationMethod
	if useIsolation {
		application = ApplyFirstOnly
	} else {
		application = ApplyAllCombined
	}

	return &AlgorithmContext{applicationMethod: application}
}
