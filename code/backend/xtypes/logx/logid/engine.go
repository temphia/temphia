package logid

const (
	BinderEventProcessStart   = "binder_event_process_start"
	BinderExecutePanicked     = "binder_execute_panicked"
	BinderEventRequestDebug   = "binder_event_request_debug"
	BinderEventResponseDebug  = "binder_event_response_debug"
	BinderExecuteErr          = "binder_execute_err"
	BinderEventProcessOK      = "binder_event_process_ok"
	BinderExecutionLog        = "binder_execution_log"
	BinderExecutionLogBatched = "binder_execution_log_batched"
	BinderExecutionForked     = "binder_execution_self_fork"
)

const (
	EngineServePlugLoadError           = "engine_serve_plug_load_err"
	EngineServeAgentLoadError          = "engine_serve_agent_load_err"
	EngineServeEmptyMappings           = "engine_serve_empty_mappings"
	EngineServeBprintErr               = "engine_serve_bprint_err"
	EngineExecServeAgentLoadError      = "engine_exec_serve_agent_load_err"
	EngineExecServeExecBuilderNotFound = "engine_exec_serve_builder_not_found"
	EngineExecServeExecBuilderErr      = "engine_exec_serve_builder_err"
	EngineExecAction                   = "engine_exec_action"
	EngineResourcesLoading             = "engine_exec_resource_loading"
	EngineResourcesLoaded              = "engine_exec_resource_loaded"
)

// controller

const (
	EngineStartupHookLoadErr    = "startup_hook_load_err"
	EngineStartupHookExecuteErr = "startup_hook_execute_err"
)
