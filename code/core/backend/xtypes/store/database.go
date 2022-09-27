package store

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type CoreHub interface {
	CoreDB

	Inject(app xtypes.App)
}

type SyncDB interface {
	RepoOps
	BprintOps
	PlugOps
	AgentOps
}

type CoreDB interface {
	SystemOps
	TenantOps
	UserOps
	SyncDB

	UserMessageOps
	UserGroupExtra
	// GetInnerDriver() any
	// Migrate() error
}

type SystemOps interface {
	AddSystemEvent(data *entities.SystemEvent) error
	RemoveSystemEvent(id int64) error
	ListSystemEvent(tenantId, etype string, last int64) ([]*entities.SystemEvent, error)

	AddSystemKV(tenantId, data *entities.SystemKV) error
	UpdateSystemKV(tenantId, key, ktype string, data map[string]any) error
	GetSystemKV(tenantId, key, ktype string) (*entities.SystemKV, error)
	RemoveSystemKV(tenantId, key, ktype string) error
	ListSystemKV(tenantId, etype string) ([]*entities.SystemKV, error)
}

type TenantOps interface {
	AddTenant(tenant *entities.Tenant) error
	UpdateTenant(slug string, data map[string]any) error
	GetTenant(tenant string) (*entities.Tenant, error)
	RemoveTenant(slug string) error
	ListTenant() ([]*entities.Tenant, error)

	AddTenantHook(data *entities.TenantHook) error
	UpdateTenantHook(tenantId, targetType string, id int64, data map[string]any) error
	ListTenantHook(tenantId, targetType string) ([]*entities.TenantHook, error)
	GetTenantHook(tenantId, targetType string, id int64) (*entities.TenantHook, error)
	RemoveTenantHook(tenantId, targetType string, id int64) error

	AddDomain(domain *entities.TenantDomain) error
	UpdateDomain(tenantId string, id int64, data map[string]any) error
	GetDomain(tenantId string, id int64) (*entities.TenantDomain, error)
	GetDomainByName(tenantId string, name string) (*entities.TenantDomain, error)

	RemoveDomain(tenantId string, id int64) error
	ListDomain(tenantId string) ([]*entities.TenantDomain, error)
}

type UserOps interface {
	AddUserGroup(ug *entities.UserGroup) error
	GetUserGroup(tenantId string, slug string) (*entities.UserGroup, error)
	ListUserGroups(tenantId string) ([]*entities.UserGroup, error)
	UpdateUserGroup(tenantId, slug string, data map[string]any) error
	RemoveUserGroup(tenantId string, ugslug string) error

	AddUser(user *entities.User, data *entities.UserData) error
	UpdateUser(tenantId, user string, data map[string]any) error
	RemoveUser(tenantId string, username string) error
	GetUserByID(tenantId string, username string) (*entities.User, error)
	GetUserByEmail(tenantId string, email string) (*entities.User, error)
	ListUsers(tenantId string) ([]*entities.User, error)
	ListUsersByGroup(tenantId, group string) ([]*entities.User, error)
	ListUsersMulti(tenantId string, users ...string) ([]*entities.User, error)

	GetUserData(tenantId string, slug string) (*entities.UserData, error)
	UpdateUserData(tenantId, slug string, data map[string]any) error
}

type UserPermissionOps interface {
	AddPerm(data *entities.Permission) error
	UpdatePerm(data map[string]any) error
	GetPerm(tenantId string, id int64) (*entities.Permission, error)
	RemovePerm(tenantId string, id int64) error

	AddRole(data *entities.Role) error
	GetRole(tenantId string, id int64) (*entities.Role, error)
	UpdateRole(data map[string]any) error
	RemoveRole(data *entities.Role) error
	AddUserRole(data *entities.UserRole) error
	GetUserRole(tenantId string, id int64) (*entities.UserRole, error)
	UpdateUserRole(data map[string]any) error
	RemoveUserRole(data *entities.UserRole) error
	ListAllPerm(tenantId string) ([]*entities.Permission, error)
	ListAllRole(tenantId string) ([]*entities.Permission, error)
	ListAllUserRole(tenantId string) ([]*entities.Permission, error)
	ListAllUserPerm(tenantId string) ([]*entities.Permission, error)
	ListUserPerm(tenantId string, userId, objType, objsub string) ([]*entities.Permission, error)
}

type UserGroupExtra interface {
	AddUserGroupAuth(data *entities.UserGroupAuth) error
	UpdateUserGroupAuth(tenantId string, gslug string, id int64, data map[string]any) error
	ListUserGroupAuth(tenantId string, gslug string) ([]*entities.UserGroupAuth, error)
	GetUserGroupAuth(tenantId string, gslug string, id int64) (*entities.UserGroupAuth, error)
	RemoveUserGroupAuth(tenantId, gslug string, id int64) error

	AddUserGroupPlug(data *entities.UserGroupPlug) error
	UpdateUserGroupPlug(tenantId string, gslug string, id int64, data map[string]any) error
	ListUserGroupPlug(tenantId string, gslug string) ([]*entities.UserGroupPlug, error)
	GetUserGroupPlug(tenantId string, gslug string, id int64) (*entities.UserGroupPlug, error)
	RemoveUserGroupPlug(tenantId, gslug string, id int64) error

	AddUserGroupData(data *entities.UserGroupData) error
	UpdateUserGroupData(tenantId string, gslug string, id int64, data map[string]any) error
	ListUserGroupData(tenantId string, gslug string) ([]*entities.UserGroupData, error)
	GetUserGroupData(tenantId string, gslug string, id int64) (*entities.UserGroupData, error)
	RemoveUserGroupData(tenantId, gslug string, id int64) error
}

type UserMessageOps interface {
	AddUserMessage(msg *entities.UserMessage) (int64, error)
	UserMessageSetRead(tenantId, user string, id int64) error
	RemoveUserMessage(tenantId string, username string, id int64) error
	ListUserMessages(tenantId string, data *entities.UserMessageReq) ([]*entities.UserMessage, error)

	ReadUserMessages(tenantId, userId string, id []int64) error
	DeleteUserMessages(tenantId, userId string, id []int64) error
}

// sync_db

type BprintOps interface {
	BprintNew(tenantId string, et *entities.BPrint) error
	BprintUpdate(tenantId, id string, data map[string]any) error
	BprintGet(tenantId, id string) (*entities.BPrint, error)
	BprintDel(tenantId, id string) error
	BprintList(tenantId, group string) ([]*entities.BPrint, error)
}

type PlugOps interface {
	PlugNew(tenantId string, pg *entities.Plug) error
	PlugUpdate(tenantId string, id string, data map[string]any) error
	PlugGet(tenantId, pid string) (*entities.Plug, error)
	PlugDel(tenantId, pid string) error
	PlugList(tenantId string) ([]*entities.Plug, error)

	PlugListByBprint(tenantId, bprint string) ([]*entities.Plug, error)

	AgentNew(tenantId string, data *entities.Agent) error
	AgentUpdate(tenantId, pid, id string, data map[string]any) error
	AgentGet(tenantId, pid, id string) (*entities.Agent, error)
	AgentDel(tenantId, pid, agentId string) error
	AgentList(tenantId, pid string) ([]*entities.Agent, error)

	ResourceNew(tenantId string, obj *entities.Resource) error
	ResourceUpdate(tenantId string, id string, data map[string]any) error
	ResourceGet(tenantId, rid string) (*entities.Resource, error)
	ResourceDel(tenantId, rid string) error
	ResourceList(tenantId string) ([]*entities.Resource, error)
	ResourcesMulti(tenantId string, rids ...string) ([]*entities.Resource, error)
	ResourcesByTarget(tenantId string, target string) ([]*entities.Resource, error)
	ResourceListByPlug(tenantId string, plugId string) ([]*entities.Resource, error)
}

type AgentOps interface {
	AgentLinkNew(tenantId string, data *entities.AgentLink) error
	AgentLinkUpdate(tenantId, pid, aid string, id int64, data map[string]any) error
	AgentLinkGet(tenantId, pid, aid string, id int64) (*entities.AgentLink, error)
	AgentLinkDel(tenantId, pid, aid string, id int64) error
	AgentLinkList(tenantId, pid, aid string) ([]*entities.AgentLink, error)
	AgentLinkListReverse(tenantId, pid, aid string) ([]*entities.AgentLink, error)

	AgentExtensionNew(tenantId string, data *entities.AgentExtension) error
	AgentExtensionUpdate(tenantId, pid, aid string, id int64, data map[string]any) error
	AgentExtensionGet(tenantId, pid, aid string, id int64) (*entities.AgentExtension, error)
	AgentExtensionDel(tenantId, pid, aid string, id int64) error
	AgentExtensionList(tenantId, pid, aid string) ([]*entities.AgentExtension, error)

	AgentResourceNew(tenantId string, data *entities.AgentResource) error
	AgentResourceUpdate(tenantId, pid, aid, slug string, data map[string]any) error
	AgentResourceGet(tenantId, pid, aid, slug string) (*entities.AgentResource, error)
	AgentResourceDel(tenantId, pid, aid, slug string) error
	AgentResourceList(tenantId, pid, aid string) ([]*entities.AgentResource, error)

	ResourceListByAgent(tenantId string, pid, aid string) ([]*entities.Resource, error)
}

type RepoOps interface {
	RepoNew(tenantId string, pg *entities.Repo) error
	RepoUpdate(tenantId string, id int64, data map[string]any) error
	RepoGet(tenantId string, id int64) (*entities.Repo, error)
	RepoDel(tenantId string, id int64) error
	RepoList(tenantId string) ([]*entities.Repo, error)
}
