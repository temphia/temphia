package tasmsdk

/*



	ListUsers(group string) ([]string, error)
	MessageUser(group, user, message string, encrypted bool) error
	GetUser(group, user string) (*entities.UserInfo, error)

	MessageCurrentUser(user, message string, encrypted bool) error
	CurrentUser() (*entities.UserInfo, error)

	HttpRaw(*HTTPRequest) *HTTPResponse
	HttpBatch([]HTTPRequest) []HTTPResponse


	HttpQuickGet(url string) ([]byte, error)
	HttpQuickPost(url string, data []byte) ([]byte, error)

	HttpFormPost(url string, headers []HttpHeaderPair, data []byte) ([]byte, error)
	HttpJsonGet(url string, headers []HttpHeaderPair) ([]byte, error)
	HttpJsonPost(url string, headers []HttpHeaderPair, data []byte) ([]byte, error)

	SelfGetFile(file string) ([]byte, error)
	SelfAddFile(file string, data []byte) error
	SelfUpdateFile(file string, data []byte) error

	SelfListResources() ([]*Resource, error)
	SelfGetResource(name string) (*Resource, error)

	SelfInLinks() ([]Link, error)
	SelfOutLinks() ([]Link, error)

	SelfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error)
	SelfModuleExec(name, method, path string, data xtypes.LazyData) (xtypes.LazyData, error)
	SelfInvokerExec(method string, data xtypes.LazyData) (xtypes.LazyData, error)
*/
