package snake

import (
	"testing"
)

func makeValidGene() Gene {
	return Gene{
		Trait: Albino,
		TraitRelations: []TraitRelation{{
			Allele:     Albino,
			Expression: Recessive,
		}},
		Inheritance: Mendelian,
		Description: "albino",
	}
}
func TestGene(t *testing.T) {

	g := makeValidGene()

	err := g.Validate()
	if err != nil {
		t.Error("expected valid gene when all fields are properly set")
	}

	g = makeValidGene()
	g.TraitRelations = []TraitRelation{}
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when AllelicRelations is empty")
	}

	g = makeValidGene()
	g.Description = ""
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when description is empty")
	}

}
