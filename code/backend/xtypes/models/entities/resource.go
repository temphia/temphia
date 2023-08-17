package entities

import (
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

type Resource struct {
	Id          string     `json:"id,omitempty" db:"id,omitempty"`
	Name        string     `json:"name,omitempty" db:"name,omitempty"`
	Type        string     `json:"type,omitempty" db:"type,omitempty"`
	SubType     string     `json:"sub_type,omitempty" db:"sub_type,omitempty"`
	Target      string     `json:"target,omitempty" db:"target,omitempty"`
	Payload     string     `json:"payload,omitempty" db:"payload,omitempty"`
	Policy      string     `json:"policy,omitempty" db:"policy,omitempty"`
	OwnedByPlug string     `json:"owned_by_plug,omitempty"  db:"owned_by_plug,omitempty"`
	ExtraMeta   JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId    string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

func (r *Resource) SplitTarget(expected int) ([]string, error) {
	targets := strings.Split(r.Target, "/")

	if len(targets) != expected {
		return nil, easyerr.Error("could not parse target")
	}
	return targets, nil
}

// ResourcePair is container with resource and AgentResource
type ResourcePair struct {
	AgentResource *AgentResource
	Resource      *Resource
}
