package retic

import (
	"time"

	"github.com/google/uuid"
)

type Dominance uint8

const (
	Dominant Dominance = iota
	Recessive
)

type Expression uint8

const (
	CompleteDominance Expression = iota
	IncompleteDominance
	Codominance
	SexLinked
)

type AllelePair struct {
	A1 Dominance
	A2 Dominance
}

type Trait struct {
	Name        string
	Description string
	Expression  Expression
	AllelePair
}

type Retic struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Birth      time.Time `json:"birth"`
	Genus      string    `json:"genus"`
	Species    string    `json:"species"`
	Subspecies string    `json:"subspecies"`
	SireId     string    `json:"sire_id"`
	DameId     string    `json:"dame_id"`
	RoomId     string    `json:"room_id"`
	RackId     string    `json:"rack_id"`
	BinId      string    `json:"bin_id"`
	Traits     []Trait   `json:"traits"`
}

func PredictPairing(r1 Retic, r2 Retic) (map[string][]AllelePair, error) {
	m := make(map[string][]AllelePair)

	for _, t1 := range r1.Traits {
		for _, t2 := range r2.Traits {
			if t1.Name == t2.Name {
				m[t1.Name] = append(m[t1.Name],
					AllelePair{A1: t1.A1, A2: t2.A1},
					AllelePair{A1: t1.A1, A2: t2.A2},
					AllelePair{A1: t1.A2, A2: t2.A1},
					AllelePair{A1: t1.A2, A2: t2.A2},
				)
			}
		}
	}

	return m, nil
}

// Can only pass forward one allele
//   Z  Z
// Z ZZ ZZ
// W WZ WZ
