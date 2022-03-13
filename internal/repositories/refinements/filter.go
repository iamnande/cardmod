package refinements

// Filter is used to filter refinement collection results.
type Filter struct {
	source string
	target string
}

func NewFilter(source, target string) *Filter {
	return &Filter{
		source: source,
		target: target,
	}
}

func (f Filter) Source() string {
	return f.source
}

func (f Filter) Target() string {
	return f.target
}
