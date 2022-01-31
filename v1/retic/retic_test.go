package retic

import (
	"testing"
)

func TestAllelicCombinations(t *testing.T) {
	rSingle := Retic{
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
		},
	}
	e0, e1 := "SEX", "sex"
	c := rSingle.AllelicCombinations()

	if c[0].String() != e0 {
		t.Errorf("expceted 'SEX', got '%v'", c[0])
	}
	if c[1].String() != e1 {
		t.Errorf("expected 'sex', got '%v'", c[1])
	}

	rDouble := Retic{
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
		},
	}
	e0, e1 = "SEX+GOLDEN_CHILD", "SEX+golden_child"
	e2, e3 := "sex+GOLDEN_CHILD", "sex+golden_child"
	c = rDouble.AllelicCombinations()

	if c[0].String() != e0 {
		t.Errorf("expected 'SEX+GOLDEN_CHILD', got '%v'", c[0])
	}
	if c[1].String() != e1 {
		t.Errorf("expected 'SEX+golden_child', got '%v'", c[1])
	}
	if c[2].String() != e2 {
		t.Errorf("expected 'sex+GOLDEN_CHILD', got '%v'", c[2])
	}
	if c[3].String() != e3 {
		t.Errorf("expected 'sex+golden_child', got '%v'", c[3])
	}
}
