package retic

import (
	"math"
	"strings"
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

func (hs HaploidSet) String() string {
	traitStrings := make([]string, 0)
	for _, a := range hs {
		if a.Variant == Dominant {
			traitStrings = append(traitStrings, strings.ToUpper(a.Trait))
		} else {
			traitStrings = append(traitStrings, strings.ToLower(a.Trait))
		}
	}
	return strings.Join(traitStrings, "+")
}

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
	cnt := len(r.Genes)
	hs := make([]HaploidSet, int(math.Pow(2, float64(cnt))))
	distributeAlleles(&hs, r.Genes, cnt, cnt)
	return hs
}

func distributeAlleles(hs *[]HaploidSet, genes []Genotype, index int, count int) {
	// If only one gene, then just copy the alleles over
	if len(genes) == 1 {
		for _, g := range genes {
			for i, a := range g {
				(*hs)[count-index] = append((*hs)[i], a)
			}
		}
		return
	}
	// grab the first two genes
	if len(genes) == 2 && count == 2 {
		for _, g := range genes[0:2] {
			for i, a := range g {
				(*hs)[count-index] = append((*hs)[i], a)
			}
		}
	}
	index -= 2
	if index > 0 {
		distributeAlleles(hs, genes[2:], index, count)
	}
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
