package runner

type BootstrapContext struct {
	TenantId string
	PlugId   string
	AgentId  string
	File     string
	GetFile  func() ([]byte, error)
}
