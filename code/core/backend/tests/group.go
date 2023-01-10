package tests

type Group struct {
	handle *TestHandle
	name   string
	cases  map[string]func(*TestHandle) error
}

func NewGroup(name string, handle *TestHandle) *Group {
	return &Group{
		name:   name,
		handle: handle,
	}
}

func (g *Group) AddCase(name string, fn func(*TestHandle) error) {
	g.cases[name] = fn
}

func (g *Group) Name() string { return g.name }
func (g *Group) Cases() []string {
	keys := make([]string, 0, len(g.cases))
	for k := range g.cases {
		keys = append(keys, k)
	}
	return keys
}

func (g *Group) Run(cse string) error {
	return g.cases[cse](g.handle)
}
