package bind2ipc

import (
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

type Bind2IPC struct {
	binding bindx.Core
}

func New(binding bindx.Core) Bind2IPC {
	return Bind2IPC{
		binding: binding,
	}
}

type logRequest struct {
	message string
}

func (b *Bind2IPC) Log(opts *logRequest) {
	b.binding.Log(opts.message)
}

type lazyLogRequest struct {
	Messages []string
}

func (b *Bind2IPC) LazyLog(opts *lazyLogRequest) {
	b.binding.LazyLog(opts.Messages)
}

type getFileRequest struct {
	File string
}

type getFileResponse struct {
	Data    []byte
	Version int64
}

func (b *Bind2IPC) GetFileWithMeta(opts *getFileRequest) (*getFileResponse, error) {

	out, ver, err := b.binding.GetFileWithMeta(opts.File)
	if err != nil {
		return nil, err
	}

	return &getFileResponse{
		Data:    out,
		Version: ver,
	}, nil
}

func (b *Bind2IPC) ListResources() ([]*bindx.Resource, error) {
	return b.binding.ListResources()
}

type getResourceRequest struct {
	Name string
}

func (b *Bind2IPC) GetResource(opts *getResourceRequest) (*bindx.Resource, error) {
	return b.binding.GetResource(opts.Name)
}

func (b *Bind2IPC) InLinks() ([]bindx.Link, error) {
	return b.binding.InLinks()
}

func (b *Bind2IPC) OutLinks() ([]bindx.Link, error) {
	return b.binding.OutLinks()
}

type linkExecRequest struct {
	Name   string
	Method string
	Data   []byte
}

type linkExecResponse struct {
	Data []byte
}

func (b *Bind2IPC) LinkExec(opts *linkExecRequest) (*linkExecResponse, error) {

	resp, err := b.binding.LinkExec(opts.Name, opts.Method, lazydata.NewJsonData(opts.Data))
	if err != nil {
		return nil, err
	}

	out, err := resp.AsJsonBytes()
	if err != nil {
		return nil, err
	}

	return &linkExecResponse{
		Data: out,
	}, nil
}

func (b *Bind2IPC) LinkExecEmit(opts *linkExecRequest) error {
	return b.binding.LinkExecEmit(opts.Name, opts.Method, lazydata.NewJsonData(opts.Data))
}

type newModuleRequest struct {
	Name string
	Data []byte
}

func (b *Bind2IPC) NewModule(opts *newModuleRequest) (int32, error) {
	return b.binding.NewModule(opts.Name, lazydata.NewJsonData(opts.Data))
}

type moduleTicketRequest struct {
	Name string
	Data []byte
}

func (b *Bind2IPC) ModuleTicket(opts *moduleTicketRequest) (string, error) {
	return b.binding.ModuleTicket(opts.Name, lazydata.NewJsonData(opts.Data))
}

type moduleExecRequest struct {
	MID    int32
	Method string
	Data   []byte
}

type moduleExecResponse struct {
	Data []byte
}

func (b *Bind2IPC) ModuleExec(opts *moduleExecRequest) (*moduleExecResponse, error) {
	resp, err := b.binding.ModuleExec(opts.MID, opts.Method, lazydata.NewJsonData(opts.Data))
	if err != nil {
		return nil, err
	}

	out, err := resp.AsJsonBytes()
	if err != nil {
		return nil, err
	}

	return &moduleExecResponse{
		Data: out,
	}, nil

}

type forkExecRequest struct {
	Method string
	Data   []byte
}

func (b *Bind2IPC) ForkExec(opts *forkExecRequest) error {
	return b.binding.ForkExec(opts.Method, opts.Data)
}
