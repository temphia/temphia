package job

import (
	"sync"

	"github.com/temphia/temphia/code/core/backend/xtypes/enginex/event"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

type Job struct {
	PlugId      string
	AgentId     string
	EventId     string
	EventAction string
	Namespace   string
	Payload     []byte

	PendingPrePolicy  bool
	PendingPostPolicy bool

	Loaded bool
	// lazy loadable
	Plug    *entities.Plug
	Agent   *entities.Agent
	Invoker Invoker
}

func (j *Job) AsEvent() *event.Request {
	return &event.Request{
		Id:   j.EventId,
		Name: j.EventAction,
		Data: j.Payload,
	}

}

// usage
// j.Add()
// runtime.PreformAsync(j)
// j.Wait()

type AsyncJob struct {
	Inner     Job
	result    *event.Response
	resultErr error
	wg        sync.WaitGroup
}

func (j *AsyncJob) Err(err error) {
	j.resultErr = err
	j.wg.Done()
}

func (j *AsyncJob) Ok(resp *event.Response) {
	j.result = resp
	j.wg.Done()
}

func (j *AsyncJob) Add() {
	j.wg.Add(1)
}

func (j *AsyncJob) Wait() {
	j.wg.Wait()
}

func (j *AsyncJob) Result() (*event.Response, error) {
	return j.result, j.resultErr
}
