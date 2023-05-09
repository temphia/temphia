package scoper

type GroupScoper struct {
	scopes map[string]struct{}
}

func (g *GroupScoper) HasScope(scope string) bool {
	_, ok := g.scopes[scope]
	return ok
}
