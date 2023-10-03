package enginehub

import (
	"sync"
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (e *EngineHub) runStartupHooks(tenants []string, minwait time.Duration) {

	var wg sync.WaitGroup

	for _, tenant := range tenants {

		wg.Add(1)
		go func(tenantId string) {
			defer wg.Done()

			hooks, err := e.corehub.ListTargetHookByType(tenantId, entities.TargetHookTypeStartUp, "app")
			if err != nil {
				e.logger.
					Err(err).
					Str("tenant_id", tenantId).Msg(logid.EngineStartupHookLoadErr)

				return
			}

			for _, hook := range hooks {
				go func(hook *entities.TargetHook) {
					_, err := e.engine.RPXecute(etypes.RPXecuteOptions{
						TenantId: tenantId,
						PlugId:   hook.PlugId,
						AgentId:  hook.AgentId,
						Action:   hook.Handler,
						Payload:  []byte(`{}`),
						Invoker:  nil,
					})

					if err != nil {
						e.logger.
							Err(err).
							Str("tenant_id", tenantId).
							Interface("hook", hook).
							Msg(logid.EngineStartupHookExecuteErr)

					}

				}(hook)

				time.Sleep(minwait)
			}

		}(tenant)
	}

	wg.Wait()
}
