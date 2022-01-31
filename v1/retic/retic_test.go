package retic

import (
	"fmt"
	"testing"
)

func TestAllelicCombinations(t *testing.T) {
	r := Retic{
		Genes: []Genotype{
			{
				Allele{
					Trait:    "sex",
					Behavior: CompleteDominance,
					Variant:  Dominant,
				},
				Allele{
					Trait:    "sex",
					Behavior: CompleteDominance,
					Variant:  Recessive,
				},
			},
			{
				Allele{
					Trait:    "golden_child",
					Behavior: CompleteDominance,
					Variant:  Dominant,
				},
				Allele{
					Trait:    "golden_child",
					Behavior: CompleteDominance,
					Variant:  Recessive,
				},
			},
			{
				Allele{
					Trait:    "albino",
					Behavior: CompleteDominance,
					Variant:  Dominant,
				},
				Allele{
					Trait:    "albino",
					Behavior: CompleteDominance,
					Variant:  Recessive,
				},
			},
		},
	}
	c := r.AllelicCombinations()
	fmt.Println(c)
	t.Errorf("not fully implemented")
}
