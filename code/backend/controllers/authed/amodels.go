package authed

/*

	alt_auth => generate -> { first_next -> second_next }  -> submit /hook/-> finish
	login => next -> submit /hook/-> finish
	signup => next -> submit /hook/-> finish

*/

type ListAuthResponse struct {
	PasswordAuth   bool            `json:"pass_auth,omitempty"`
	OpenSignUp     bool            `json:"open_signup,omitempty"`
	AltAuthMethods []AltAuthMethod `json:"alt_auth_method,omitempty"`
}

type AltAuthMethod struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type AuthGenerateRequest struct {
	Id        int64  `json:"-"`
	SiteToken string `json:"site_token,omitempty"`
	UserGroup string `json:"user_group,omitempty"`
}

type AuthGenerateResponse struct {
	StateToken string   `json:"state_token,omitempty"`
	AuthURL    string   `json:"auth_url,omitempty"`
	Scopes     []string `json:"scopes,omitempty"`
	ClientId   string   `json:"client_id,omitempty"`
}

type AuthNextFirstRequest struct {
	Id        int64  `json:"-"`
	AuthCode  string `json:"auth_code,omitempty"`
	AuthState string `json:"auth_state,omitempty"`
	SiteToken string `json:"site_token,omitempty"`
	UserGroup string `json:"user_group,omitempty"`
}

type AuthNextFirstResponse struct {
	Message     string   `json:"message,omitempty"`
	Ok          bool     `json:"ok,omitempty"`
	FirstToken  string   `json:"first_token,omitempty"`
	NewUser     bool     `json:"new_user,omitempty"`
	Email       string   `json:"email,omitempty"`
	UserIdHints []string `json:"user_id_hints,omitempty"`
}

type AuthNextSecondRequest struct {
	SiteToken  string `json:"site_token,omitempty"`
	FirstToken string `json:"first_token,omitempty"`
	SignUpdata struct {
		UserId   string `json:"user_id"`
		Bio      string `json:"bio"`
		FullName string `json:"full_name"`
		Profile  string `json:"profile"`
	} `json:"signup_data,omitempty"`
}

type AuthNextSecondResponse struct {
	Ok          bool   `json:"ok,omitempty"`
	Message     string `json:"message,omitempty"`
	NextToken   string `json:"next_token,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	EmailVerify bool   `json:"email_verify,omitempty"`
}

type AuthSubmitRequest struct {
	NextToken string `json:"next_token,omitempty"`
	SiteToken string `json:"site_token,omitempty"`
}

type AuthSubmitResponse struct {
	SubmitResponse
}

type LoginNextRequest struct {
	UserIdent string `json:"user_ident,omitempty"`
	Password  string `json:"password,omitempty"`
}

type LoginNextResponse struct {
	Ok          bool   `json:"ok,omitempty"`
	Message     string `json:"message,omitempty"`
	NextToken   string `json:"next_token,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	EmailVerify bool   `json:"email_verify,omitempty"`
	PassChange  bool   `json:"pass_change,omitempty"`
}

type LoginSubmitRequest struct {
	NextToken string `json:"next_token,omitempty"`
}

type LoginSubmitResponse struct {
	SubmitResponse
}

type AuthFinishRequest struct {
	PreAuthedToken string `json:"preauthed_token,omitempty"`
	ProofToken     string `json:"proof_token,omitempty"`
}

type AuthFinishResponse struct {
	UserToken string `json:"user_token,omitempty"`
	TenantId  string `json:"tenant_id,omitempty"`
}

type SubmitResponse struct {
	Message        string `json:"message,omitempty"`
	Ok             bool   `json:"ok,omitempty"`
	PreAuthedToken string `json:"preauthed_token,omitempty"`
	HasExecHook    bool   `json:"has_exec_hook,omitempty"`
	HookPlugId     string `json:"hook_plug_id,omitempty"`
	HookAgentId    string `json:"hook_agent_id,omitempty"`
	HookExecToken  string `json:"hook_exec_token,omitempty"`
}
