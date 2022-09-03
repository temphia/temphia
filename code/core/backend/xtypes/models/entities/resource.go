package entities

import (
	"strings"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities/resource"
)

type Resource struct {
	Id        string     `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Type      string     `json:"type,omitempty" db:"type,omitempty"`
	SubType   string     `json:"sub_type,omitempty" db:"sub_type,omitempty"`
	Target    string     `json:"target,omitempty" db:"target,omitempty"`
	Payload   string     `json:"payload,omitempty" db:"payload,omitempty"`
	Policy    string     `json:"policy,omitempty" db:"policy,omitempty"`
	PlugId    string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

func (r *Resource) splitTarget(expected int) ([]string, error) {
	targets := strings.Split(r.Payload, "/")
	if len(targets) != expected {
		return nil, easyerr.Error("could not parse target")
	}
	return targets, nil
}

type ResourceSockRoom struct {
	Type   string `json:"type,omitempty"`
	Value  string `json:"value,omitempty"`
	Policy string `json:"policy,omitempty"`
}

func (r *Resource) SockRoom() *ResourceSockRoom {
	return &ResourceSockRoom{
		Type:   r.Type,
		Value:  r.Payload,
		Policy: r.Policy,
	}
}

type ResourceDtable struct {
	Type   string `json:"type,omitempty"`
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	Table  string `json:"table,omitempty"`
	Policy string `json:"policy,omitempty"`
}

func (r *Resource) Dtable() *ResourceDtable {
	targets, err := r.splitTarget(3)
	if err != nil {
		panic(err)
	}

	return &ResourceDtable{
		Type:   resource.Dtable,
		Source: targets[0],
		Group:  targets[1],
		Table:  targets[2],
		Policy: r.Policy,
	}

}

type ResourceDgroup struct {
	Type   string `json:"type,omitempty"`
	Source string `json:"source,omitempty"`
	Group  string `json:"group,omitempty"`
	Policy string `json:"policy,omitempty"`
}

func (r *Resource) Dgroup() *ResourceDgroup {
	targets, err := r.splitTarget(2)
	if err != nil {
		panic(err)
	}
	return &ResourceDgroup{
		Type:   resource.Dgroup,
		Source: targets[0],
		Group:  targets[1],
		Policy: r.Policy,
	}
}

type ResourceFolder struct {
	Type   string `json:"type,omitempty"`
	Source string `json:"source,omitempty"`
	Folder string `json:"group,omitempty"`
	Policy string `json:"policy,omitempty"`
}

func (r *Resource) Folder() *ResourceFolder {
	targets, err := r.splitTarget(2)
	if err != nil {
		panic(err)
	}
	return &ResourceFolder{
		Type:   resource.Folder,
		Source: targets[0],
		Folder: targets[1],
		Policy: r.Policy,
	}

}
