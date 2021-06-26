package cat

type Cat struct{}

func New() *Cat {
	return &Cat{}
}

func (Cat) Say() string {
	return "Meaw"
}

func (Cat) Run() string {
	return "Run"
}
