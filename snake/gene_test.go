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

}
