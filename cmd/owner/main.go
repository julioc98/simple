package main

import (
	"simple/cat"
	"simple/dog"
	"simple/falcon"
	"simple/owner"
)

func main() {
	dogOwner := owner.New(dog.New())
	dogOwner.PetGo()

	catOwner := owner.New(cat.New())
	catOwner.PetGo()

	falconOwner := owner.New(falcon.New())
	falconOwner.PetGo()
}
