package snake

import "testing"

func makeValidGene() Gene {
	return Gene{
		Allele: Albino,
		AllelicRelations: []AllelicRelation{{
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
	g.AllelicRelations = []AllelicRelation{}
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when AllelicRelations is empty")
	}

	g = makeValidGene()
	g.Percentage = 49
	g.Inheritance = Mendelian
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when percentage is not 50 or 100 and inheritance is Mendelian")
	}

	g = makeValidGene()
	g.Percentage = 101
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when percentage is greater than 100")
	}

	g = makeValidGene()
	g.Percentage = 99
	g.Inheritance = Mendelian
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when percentage is not 50 or 100 and inheritance is Mendelian")
	}

	g = makeValidGene()
	g.Percentage = -1
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when percentage is less than 0")
	}

	g = makeValidGene()
	g.Description = ""
	err = g.Validate()
	if err == nil {
		t.Error("expected errors when description is empty")
	}

}
