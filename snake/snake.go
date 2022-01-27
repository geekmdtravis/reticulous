package snake

import (
	"time"

	"github.com/geekmdtravis/reticulous/blockchain"
	"github.com/geekmdtravis/reticulous/taxonomy"
)

type Snake struct {
	DateOfBirth time.Time
	Name        string
	NFTAddress  string
	NFTNetwork  blockchain.Network
	Gender      Gender
	Centimeters uint8
	Grams       uint8
	Sire        *Snake
	Dam         *Snake
	Genes       []Gene
	Locales     map[Trait]float32
	taxonomy.Taxonomy
}
