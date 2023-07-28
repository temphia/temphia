package service

type DyndbLock interface {
	GlobalLock(tenantId string) (string, error)
	GlobalUnLock(tenantId, utoken string) error
	GroupLock(tenantId, group string) (string, error)
	GroupUnLock(tenantId, group, utoken string) error
}
