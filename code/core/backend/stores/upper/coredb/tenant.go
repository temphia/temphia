package coredb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

// tenant

func (d *DB) AddTenant(tenant *entities.Tenant) error {
	_, err := d.tenantTable().Insert(tenant)
	return err
}

func (d *DB) UpdateTenant(slug string, data map[string]interface{}) error {
	return d.tenantTable().Find(db.Cond{"slug": slug}).Update(data)
}

func (d *DB) GetTenant(slug string) (*entities.Tenant, error) {
	ten := &entities.Tenant{}
	err := d.tenantTable().Find(db.Cond{"slug": slug}).One(ten)
	if err != nil {
		return nil, err
	}
	return ten, nil
}

func (d *DB) RemoveTenant(slug string) error {
	return d.tenantTable().Find(db.Cond{"slug": slug}).Delete()
}

func (d *DB) ListTenant() ([]*entities.Tenant, error) {
	tens := make([]*entities.Tenant, 0)
	err := d.tenantTable().Find().All(&tens)
	if err != nil {
		return nil, err
	}

	return tens, nil
}

// hooks

func (d *DB) AddTenantHook(data *entities.TenantHook) error {
	_, err := d.tenantHookTable().Insert(data)
	return err
}

func (d *DB) UpdateTenantHook(tenantId string, target string, id int64, data map[string]any) error {
	return d.tenantHookTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Update(data)
}

func (d *DB) GetTenantHook(tenantId string, target string, id int64) (*entities.TenantHook, error) {
	w := &entities.TenantHook{}
	err := d.tenantHookTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).One(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (d *DB) RemoveTenantHook(tenantId, target string, id int64) error {
	return d.tenantHookTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Delete()
}

func (d *DB) ListTenantHook(tenantId string, target string) ([]*entities.TenantHook, error) {
	ws := make([]*entities.TenantHook, 0)

	cond := db.Cond{"tenant_id": tenantId}
	if target != "" {
		cond["target"] = target
	}

	err := d.tenantHookTable().Find().All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

// domain

func (d *DB) AddDomain(domain *entities.TenantDomain) error {
	_, err := d.tenantDomainTable().Insert(domain)
	return err
}

func (d *DB) UpdateDomain(tenantId string, id int64, data map[string]interface{}) error {
	return d.tenantDomainTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Update(data)
}

func (d *DB) GetDomain(tenantId string, id int64) (*entities.TenantDomain, error) {
	td := &entities.TenantDomain{}

	err := d.tenantDomainTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).One(td)
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (d *DB) GetDomainByName(tenantId string, name string) (*entities.TenantDomain, error) {
	td := &entities.TenantDomain{}

	err := d.tenantDomainTable().Find(db.Cond{"name": name, "tenant_id": tenantId}).One(td)
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (d *DB) RemoveDomain(tenantId string, id int64) error {
	return d.tenantDomainTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Delete()
}

func (d *DB) ListDomain(tenantId string) ([]*entities.TenantDomain, error) {
	domains := make([]*entities.TenantDomain, 0)
	err := d.tenantDomainTable().Find(db.Cond{"tenant_id": tenantId}).All(&domains)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

// widgets

func (d *DB) AddDomainWidget(domain *entities.DomainWidget) error {
	_, err := d.domainWidgetTable().Insert(domain)
	return err
}

func (d *DB) UpdateDomainWidget(tenantId string, id int64, data map[string]interface{}) error {
	return d.domainWidgetTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Update(data)
}

func (d *DB) GetDomainWidget(tenantId string, id int64) (*entities.DomainWidget, error) {
	w := &entities.DomainWidget{}
	err := d.domainWidgetTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).One(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (d *DB) RemoveDomainWidget(tenantId string, id int64) error {
	return d.domainWidgetTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Delete()
}

func (d *DB) ListDomainWidget(tenantId string, did int64) ([]*entities.DomainWidget, error) {
	ws := make([]*entities.DomainWidget, 0)
	err := d.domainWidgetTable().Find(db.Cond{"tenant_id": tenantId, "domain_id": did}).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

// private

func (d *DB) tenantTable() db.Collection {
	return d.table("tenants")
}

func (d *DB) tenantHookTable() db.Collection {
	return d.table("tenant_hooks")
}

func (d *DB) tenantDomainTable() db.Collection {
	return d.table("tenant_domains")
}

func (d *DB) domainWidgetTable() db.Collection {
	return d.table("domain_widgets")
}
