package corehub

type TenantStat struct {
	NoOfUsers     string `json:"no_of_users,omitempty"`
	NoOfPlugs     string `json:"no_of_plugs,omitempty"`
	NoOfDataGroup string `json:"no_of_data_group,omitempty"`
}

func (c *CoreHub) GetTenantStats() {

}
