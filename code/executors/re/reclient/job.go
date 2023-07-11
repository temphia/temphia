package reclient

type Handler func(ctx Jobctx) (any, error)

type Jobctx struct {
	data []byte
}

func (j *Jobctx) Data() []byte {
	return j.data
}

func (j *Jobctx) GetBindings() any {
	return nil
}

func (j *Jobctx) RootBindings() any {

	return nil
}
