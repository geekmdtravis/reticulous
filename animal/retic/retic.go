package retic

import "github.com/geekmdtravis/reticulous/animal"

type Snake struct {
	animal.Animal
	Centimeters uint8
	Grams       uint8
	Genes       []Gene
	Traits      []string
	Sire        *Snake
	Dam         *Snake
}
