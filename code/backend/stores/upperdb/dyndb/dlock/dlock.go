package dlock

import "time"

// fixme => remove

type LockRequest struct {
	TenantId  string
	ServiceId string
	Nested    bool
	NestedKey string
	Expiry    time.Duration
}

type UnLockRequest struct {
	TenantId    string
	ServiceId   string
	Nested      bool
	NestedKey   string
	UnlockToken string
}

type Coordinator interface {
	Lock(LockRequest) (string, error)
	UnLock(UnLockRequest) error
}

type dlock struct {
	cdtr        Coordinator
	groupExpiry time.Duration
	tableExpiry time.Duration
	serviceKey  string
}

func New(source string) *dlock {
	return &dlock{
		cdtr:        NewNoop(),
		groupExpiry: time.Minute,
		tableExpiry: time.Minute * 2,
		serviceKey:  "dyndb" + source,
	}
}

func (d *dlock) GlobalLock(tenantId string) (string, error) {
	return d.cdtr.Lock(LockRequest{
		TenantId:  tenantId,
		ServiceId: d.serviceKey,
		Nested:    false,
		NestedKey: "",
		Expiry:    d.groupExpiry,
	})
}

func (d *dlock) GlobalUnLock(tenantId, utoken string) error {
	return d.cdtr.UnLock(UnLockRequest{
		TenantId:    tenantId,
		ServiceId:   d.serviceKey,
		Nested:      false,
		NestedKey:   "",
		UnlockToken: utoken,
	})
}
func (d *dlock) GroupLock(tenantId, table string) (string, error) {
	return d.cdtr.Lock(LockRequest{
		TenantId:  tenantId,
		ServiceId: d.serviceKey,
		Nested:    true,
		NestedKey: table,
		Expiry:    d.groupExpiry,
	})
}
func (d *dlock) GroupUnLock(tenantId, table, utoken string) error {
	return d.cdtr.UnLock(UnLockRequest{
		TenantId:    tenantId,
		ServiceId:   d.serviceKey,
		Nested:      true,
		NestedKey:   table,
		UnlockToken: utoken,
	})
}

type noop struct{}

func NewNoop() *noop {
	return &noop{}
}

func (c *noop) Lock(LockRequest) (string, error) {
	return "noop", nil
}

func (c *noop) UnLock(UnLockRequest) error {
	return nil
}
