package authz

import "strings"

type scoper struct {
	scopes map[string]struct{}
}

func (s *scoper) check(scope string) bool {
	_, ok := s.scopes[scope]
	return ok
}

func (s *scoper) build(scopes string) {
	for _, scope := range strings.Split(scopes, ",") {
		s.scopes[scope] = struct{}{}
	}
}
