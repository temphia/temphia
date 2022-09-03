package vmodels

type SiteData struct {
	SiteToken string `json:"site_token,omitempty"`
	PlugURL   string `json:"plug_url,omitempty"`
	ApiURL    string `json:"api_url,omitempty"`
	TenantId  string `json:"tenant_id,omitempty"`
}
