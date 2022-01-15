package main

import (
	optional "go.eth-p.dev/goptional"
)

type Pet struct {
	species string
	name    string
}

func (p Pet) Species() string {
	return p.name
}

func (p Pet) Name() string {
	return p.name
}

type Family struct {
	pets []Pet
}

func (f Family) GetPet(name string) optional.Optional[Pet] {
	for _, pet := range f.pets {
		if pet.name == name {
			return optional.New(pet)
		}
	}

	return optional.None[Pet]()
}

func main() {
	family := Family{
		pets: []Pet{
			{species: "dog", name: "Yuki"},
			{species: "bird", name: "Koji"},
		},
	}

	koji := family.GetPet("Koji")
	kojiSpecies := optional.Map(koji, Pet.Species)
	println(kojiSpecies.Expect("value present"))
}
