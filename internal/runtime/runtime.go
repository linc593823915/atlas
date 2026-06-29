package runtime

type Runtime struct {
}

func CreateRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Start() error {
	return nil
}

func (r *Runtime) Stop() error {
	return nil
}
