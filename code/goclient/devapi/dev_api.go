package devapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

type DevAPI struct {
	http  http.Client
	url   string
	token string
}

func New(url string, token string) *DevAPI {
	return &DevAPI{
		http:  http.Client{},
		url:   url,
		token: token,
	}
}

func (d *DevAPI) BprintFileList() (map[string]string, error) {
	resp, err := d.httpPerform(http.MethodGet, "/dev/bprint/file", nil, "")
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	err = unpack(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DevAPI) BprintFilePush(file string, data []byte) error {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	fwriter, err := writer.CreateFormFile("files", file)
	if err != nil {
		return err
	}

	_, err = fwriter.Write(data)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	_, err = d.httpPerform(http.MethodPost, "/dev/bprint/file", buf, writer.FormDataContentType())
	return err
}

func (d *DevAPI) BprintFileGet(file string) ([]byte, error) {

	resp, err := d.httpPerform(http.MethodGet, fmt.Sprintf("/dev/bprint/file/%s", file), nil, "")
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func (d *DevAPI) BprintFileDel(file string) error {
	_, err := d.httpPerform(http.MethodDelete, fmt.Sprintf("/dev/bprint/file/%s", file), nil, "")
	return err
}

func (d *DevAPI) ExecWatchURL(pid, aid string) string {

	_ubase := strings.Replace(d.url, "http://", "ws://", 1)
	_ubase = strings.Replace(_ubase, "https://", "wss://", 1)

	return fmt.Sprintf("%s/dev/exec/watch/plug/%s/agent/%s?token=%s", _ubase, pid, aid, d.token)
}

func (d *DevAPI) ExecReset(pid, aid string) error {
	_, err := d.httpPerform(http.MethodPost, fmt.Sprintf("/exec/reset/plug/%s/agent/%s", pid, aid), nil, "")
	return err
}

func (d *DevAPI) ExecRun(pid, aid, action string, payload any) (any, error) {
	outr, err := pack(payload)
	if err != nil {
		return nil, err
	}

	resp, err := d.httpPerform(http.MethodPost, fmt.Sprintf("/exec/run/plug/%s/agent/%s/%s", pid, aid, action), outr, "")
	if err != nil {
		return nil, err
	}

	var a any
	err = unpack(resp, &a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// func (d *DevAPI) ModifyPlug(pid, aid string, data map[string]any) error  { return nil }
// func (d *DevAPI) ModifyAgent(pid, aid string, data map[string]any) error { return nil }

// private

func (d *DevAPI) httpPerform(method, url string, pr io.Reader, ctype string) (*http.Response, error) {

	url = fmt.Sprintf("%s%s", d.url, url)

	req, err := http.NewRequest(method, url, pr)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", d.token)
	if ctype != "" {
		req.Header.Set("Content-Type", ctype)
	}

	resp, err := d.http.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, easyerr.Error(resp.Status)
	}

	return resp, nil

}

func unpack(resp *http.Response, target any) error {
	return json.NewDecoder(resp.Body).Decode(target)
}

func pack(payload any) (io.Reader, error) {
	out, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(out), nil
}
