package falcon

type Falcon struct{}

func New() *Falcon {
	return &Falcon{}
}

func (Falcon) Say() string {
	return "kak, kak, kak"
}

func (Falcon) Run() string {
	return "run"
}

func (Falcon) Fly() string {
	return "high"
}
