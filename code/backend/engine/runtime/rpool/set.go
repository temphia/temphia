package rpool

type plugSet struct {
	agents   Set
	bprintId string
}

type agentSet struct {
	binders      Set
	counter      int64
	epochVersion int64
}

func (a *agentSet) isEmpty() bool {
	return a.binders.IsEmpty()
}

type Set map[string]struct{}

func (s Set) Pop() string {
	for v := range s {
		delete(s, v)
		return v
	}
	return ""
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Push(v string) {
	s[v] = struct{}{}
}
