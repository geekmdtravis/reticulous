package retic

import (
	"fmt"
	"math"
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

type Allele struct {
	Trait    string
	Behavior Expression
	Variant  Dominance
}

type Genotype [2]Allele

type HaploidSet []Allele

type Retic struct {
	Id         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Birth      time.Time  `json:"birth"`
	Genus      string     `json:"genus"`
	Species    string     `json:"species"`
	Subspecies string     `json:"subspecies"`
	SireId     string     `json:"sire_id"`
	DameId     string     `json:"dame_id"`
	RoomId     string     `json:"room_id"`
	RackId     string     `json:"rack_id"`
	BinId      string     `json:"bin_id"`
	Genes      []Genotype `json:"genes"`
}

// Takes a set of genes and turns them into a set of allelic combinations.
// Simple example:
// Female [ZW] will turn into the haploid set [Z, W]
// Intermediate example:
// Female het Golden child [ZW, Gg] will turn into the haploid set [ZG, Zg, WG, Wg]
// Complex example:
// Female, het Golden Child, het Albino [ZW, Gg, Aa] will turn into
// The haploid set [ZGA, ZGa, ZgA, Zga, WGA, WGa, WgA, Wga]
func (r Retic) AllelicCombinations() []HaploidSet {
	geneCnt := len(r.Genes)
	combinations := int(math.Pow(2, float64(geneCnt)))
	hs := make([]HaploidSet, combinations)

	for i := 0; i < combinations; i++ {
		// From r.Genes, this indicates the index of the current gene
		geneIndex := i % geneCnt
		// Inverted gene index used for setting values in the haploid set
		invertIndex := geneCnt - geneIndex
		fmt.Println(invertIndex)
		g := r.Genes[geneIndex]
		fmt.Println(g)
	}
	return hs
}

func PredictPairing(r1 Retic, r2 Retic) ([][]Genotype, error) {

	// TODO: return error if r1 and r2 are not opposite gender

	// TODO: Get full set of possible allele pairs that can be sent for each trait
	// and for each snake.

	return nil, nil
}

// Example
// Male: ZZGg | Female: ZWGG
//     ZG   Zg   ZG   Zg
// ZG  ZZGG ZZGg ZZGG ZZGg
// ZG  ZZGG ZZGg ZZGG ZZGg
// WG  ZWGG ZWGg ZWGG ZWGg
// WG  ZWGG ZWGg ZWGG ZWGg
//
// Steps:
// 1. Get all possible combinations for M and F from their allele set.
// 2. Combine the two sets of combinations.
//
// For n alleles there are n^2 combinations

/*
Approaching the problem programmatically; take the following steps:
Z,W x G,g x A, a

// Multiply the first two
ZG, Zg, WG, Wg x A, a

// Multiply the result by the second two
ZGA, ZgA, WGA, WGA, ZGa, Zga, WGa, Wga

Question: will this require recursion?
*/
