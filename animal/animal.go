package animal

import (
	"time"

	"github.com/geekmdtravis/reticulous/blockchain"
	"github.com/geekmdtravis/reticulous/taxonomy"
)

type Animal struct {
	DateOfBirth time.Time
	Name        string
	NFTAddress  string
	NFTNetwork  blockchain.Network
	Gender      Gender
	taxonomy.Taxonomy
}

type Gender uint8

const (
	Male Gender = iota
	Female
	Intersex
	Unsexed
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	case Intersex:
		return "intersex"
	case Unsexed:
		return "unsexed"
	}
	return "unknown"
}
