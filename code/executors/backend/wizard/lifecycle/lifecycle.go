package lifecycle

type LifecCycle interface {
	Bindings() map[string]interface{}
}

const (
	SPLASH_ON_LOAD   = "on_load"
	SPLASH_ON_SUBMIT = "on_submit"

	GROUP_BEFORE_START = "before_start"
	GROUP_AFTER_START  = "after_start"
	GROUP_BEFORE_END   = "before_end"
	GROUP_AFTER_END    = "after_end"

	GROUP_BEFORE_NEXT = "before_next"
	GROUP_BEFORE_BACK = "before_back"

	STAGE_BEFORE_VERIFY   = "before_verify"
	STAGE_AFTER_VERIFY    = "after_verify"
	STAGE_BEFORE_GENERATE = "before_generate"
	STAGE_AFTER_GENERATE  = "after_generate"
)
