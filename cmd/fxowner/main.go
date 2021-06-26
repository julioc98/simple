package main

import (
	"context"
	"fmt"
	"simple/dog"
	"simple/owner"

	"go.uber.org/fx"
)

type PetOwner interface {
	PetGo()
}

func main() {
	fx.New(
		fx.Provide(
			// dog
			func(d *dog.Dog) owner.Pet { return d },
			dog.New,

			// owner
			func(o *owner.Owner) PetOwner { return o },
			owner.New,
		),
		fx.Invoke(register),
	).Run()
}

func register(lifecycle fx.Lifecycle, petOwner PetOwner) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Hi")
				petOwner.PetGo()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				petOwner.PetGo()
				fmt.Println("Bye")
				return nil
			},
		},
	)
}
