package dog

type Dog struct{}

func New() *Dog {
	return &Dog{}
}

func (Dog) Say() string {
	return "Au Au"
}

func (Dog) Run() string {
	return "Runnn"
}
