type Instance struct {
	ID    string
	Type  string
	State string
}

func (i *Instance) Start() {
	i.State = "running"
}

func (i *Instance) Stop() {
	i.State = "stopped"
}