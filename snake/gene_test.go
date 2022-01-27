package snake

import (
	"testing"
)

func makeValidGene() Gene {
	return Gene{
		Trait: Albino,
		TraitRelations: []TraitRelation{{
			Trait:      Albino,
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

	if g.Inheritance() != Mendelian {
		t.Error("expected Mendelian inheritance")
	}

	g.Trait = WSexGene
	if g.Inheritance() != SexLinked {
		t.Error("expected SexLinkedZWFemale inheritance")
	}

	g.Trait = LocaleMainland
	if g.Inheritance() != NonMendelian {
		t.Error("expected NonMendelian inheritance")
	}

}

func TestExpression(t *testing.T) {
	if CompleteDominance.String() != "complete dominance" {
		t.Error("expected complete dominance to be expressed as complete dominance")
	}
	if IncompleteDominance.String() != "incomplete dominance" {
		t.Error("expected IncompleteDominance to be expressed as incomplete dominance")
	}
	if Codominant.String() != "codominant" {
		t.Error("expected Codominant to be expressed as codominant")
	}
	if Recessive.String() != "recessive" {
		t.Error("expected Recessive to be expressed as recessive")
	}
	if UndeclaredExpression.String() != "undeclared" {
		t.Error("expected UndeclaredExpression to be expressed as undeclared")
	}
}

func TestInheritance(t *testing.T) {
	if Mendelian.String() != "mendelian" {
		t.Error("expected Mendelian to be expressed as mendelian")
	}
	if NonMendelian.String() != "non-mendelian" {
		t.Error("expected NonMendelian to be expressed as non-mendelian")
	}
	if SexLinked.String() != "female sex linked" {
		t.Error("expected SexLinked to be expressed as female sex linked")
	}
	if UndeclaredInheritance.String() != "undeclared" {
		t.Error("expected UndeclaredInheritance to be expressed as undeclared")
	}
}

func TestNewGene(t *testing.T) {
	g := NewGene(Albino, func(t Trait) []TraitRelation {
		return []TraitRelation{{
			Trait:      GoldenChild,
			Expression: Recessive,
		}}
	})

	if g.Trait != Albino {
		t.Error("expected gene to have trait Albino")
	}

	if len(g.TraitRelations) != 1 {
		t.Error("expected gene to have 1 trait relation")
	}

	if g.TraitRelations[0].Trait != GoldenChild {
		t.Error("expected gene to have trait GoldenChild")
	}
}
