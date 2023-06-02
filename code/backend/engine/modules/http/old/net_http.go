package net

// func (n *Binding) httpRawBatch(reqs []*bindx.HttpRequest) []*bindx.HttpResponse {

// 	resp := make([]*bindx.HttpResponse, 0, len(reqs))

// 	respChan := make(chan *bindx.HttpResponse)
// 	for _, hr := range reqs {

// 		go func(req *bindx.HttpRequest) {
// 			respChan <- httpRaw(&n.hclient, req)
// 		}(hr)
// 	}
// 	i := 0
// 	for {
// 		resp = append(resp, <-respChan)
// 		i = i + 1
// 		if i <= len(reqs) {
// 			break
// 		}
// 	}

// 	return resp
// }

// func (n *Binding) httpQuickGet(url string, headers map[string]string) ([]byte, error) {

// 	resp, err := httpRequest(&n.hclient, url, http.MethodGet, headers, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.SatusCode > 100 && resp.SatusCode < 300 {
// 		return nil, errors.New(kosher.Str(resp.Body))
// 	}

// 	return resp.Body, nil
// }

// func (n *Binding) httpQuickPost(url string, headers map[string]string, data []byte) ([]byte, error) {

// 	resp, err := httpRequest(&n.hclient, url, http.MethodPost, headers, bytes.NewReader(data))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.SatusCode > 100 && resp.SatusCode < 300 {
// 		return nil, errors.New(kosher.Str(resp.Body))
// 	}

// 	return resp.Body, nil

// }

// func (n *Binding) httpFormPost(url string, headers map[string]string, data []byte) ([]byte, error) {
// 	return nil, easyerr.NotImpl()
// }

// func (n *Binding) httpJsonGet(url string, headers map[string]string) ([]byte, error) {
// 	if headers == nil {
// 		headers = StaticJsonHeader
// 	} else {
// 		headers[ContentType] = ApplicationJson
// 	}

// 	resp, err := httpRequest(&n.hclient, url, http.MethodPost, headers, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.SatusCode > 100 && resp.SatusCode < 300 {
// 		return nil, errors.New(kosher.Str(resp.Body))
// 	}

// 	return resp.Body, nil
// }

// func (n *Binding) httpJsonPost(url string, headers map[string]string, data []byte) ([]byte, error) {

// 	if headers == nil {
// 		headers = StaticJsonHeader
// 	} else {
// 		headers[ContentType] = ApplicationJson
// 	}

// 	resp, err := httpRequest(&n.hclient, url, http.MethodPost, headers, bytes.NewReader(data))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.SatusCode > 100 && resp.SatusCode < 300 {
// 		return nil, errors.New(kosher.Str(resp.Body))
// 	}

// 	return resp.Body, nil

// }
