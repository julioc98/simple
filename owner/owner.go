package owner

import (
	"fmt"
)

type Pet interface {
	Say() string
	Run() string
}

type Owner struct {
	pet Pet
}

func New(p Pet) *Owner {
	return &Owner{
		pet: p,
	}
}

func (o *Owner) PetGo() {
	fmt.Println(o.pet.Say(), o.pet.Run())
}
