package engine

import (
	"sync"
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) RunStartupHooks(tenants []string, minwait time.Duration) {

	var wg sync.WaitGroup

	for _, tenant := range tenants {

		wg.Add(1)
		go func(tenantId string) {
			defer wg.Done()

			hooks, err := c.corehub.ListTargetHookByType(tenantId, entities.TargetHookTypeStartUp, "app")
			if err != nil {
				c.logger.
					Err(err).
					Str("tenant_id", tenantId).Msg(logid.EngineStartupHookLoadErr)

				return
			}

			for _, hook := range hooks {
				go func(hook *entities.TargetHook) {
					_, err := c.engine.Execute(etypes.Execution{
						TenantId: tenantId,
						PlugId:   hook.PlugId,
						AgentId:  hook.AgentId,
						Action:   hook.Handler,
						Payload:  []byte(`{}`),
						Invoker:  nil,
					})

					if err != nil {
						c.logger.
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
