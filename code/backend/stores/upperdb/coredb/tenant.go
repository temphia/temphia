package coredb

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
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

// domain

func (d *DB) AddDomain(domain *entities.TenantDomain) (int64, error) {
	r, err := d.tenantDomainTable().Insert(domain)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
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

// private

func (d *DB) tenantTable() db.Collection {
	return d.table("tenants")
}

func (d *DB) tenantDomainTable() db.Collection {
	return d.table("domains")
}
