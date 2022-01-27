package retic

import "testing"

func makeValidGene() Gene {
	return Gene{
		Trait: Albino,
		traitRelations: []TraitRelation{{
			Allele:     Albino,
			Expression: Recessive,
		}},
		Inheritance: Mendelian,
		Zygosity:    Homozygous,
		Description: "albino",
		Percentage:  50,
	}
}
func TestGene(t *testing.T) {

	g := makeValidGene()

	err := g.Validate()
	if err != nil {
		t.Error("expected valid gene when all fields are properly set")
	}

	g = makeValidGene()
	g.traitRelations = []TraitRelation{}
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when AllelicRelations is empty")
	}

	g = makeValidGene()
	g.Percentage = 49
	g.Inheritance = Mendelian
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when percentage is not 50 and inheritance is Mendelian")
	}

	g = makeValidGene()
	g.Description = ""
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when description is empty")
	}

}
