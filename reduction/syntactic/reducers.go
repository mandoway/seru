package syntactic

type Reducer int

const (
	Perses Reducer = iota
	Vulcan
)

func (r Reducer) Name() string {
	switch r {
	case Perses:
		return "Perses"
	case Vulcan:
		return "Vulcan"
	default:
		return "Unknown"
	}
}
