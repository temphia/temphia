package vmodels

type OperatorLoginReq struct {
	User          string `json:"user,omitempty"`
	Password      string `json:"password,omitempty"`
	MasterOpToken string `json:"master_op_token,omitempty"`
}

type OperatorLoginResp struct {
	Token string `json:"token,omitempty"`
}

type NewTenant struct {
	Name          string `json:"name,omitempty"`
	Slug          string `json:"slug,omitempty"`
	SuperPassword string `json:"super_password,omitempty"`
	SuperEmail    string `json:"super_email,omitempty"`
}
