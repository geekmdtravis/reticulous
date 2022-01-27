package retic

import (
	"errors"
	"strings"

	"github.com/geekmdtravis/reticulous/snake"
)

func NewReticGene() ReticGene {
	g := ReticGene{}
	g.UseDefaultAllelicRelations()
	return g
}

type ReticGene snake.Gene

func (g ReticGene) Validate() error {
	errs := make([]string, 0, 3)

	if len(g.TraitRelations) == 0 {
		errs = append(errs, "no allelic expression relationships are defined")
	}

	if g.Inheritance == snake.Mendelian && !(int(g.Percentage) == 50) {
		errs = append(errs, "with Mendelian inheritance, percentage must be 50 as each gene can only be 1 of 2 possible")
	}

	if g.Description == "" {
		errs = append(errs, "description is empty")
	}

	if len(errs) == 0 {
		return nil
	} else {
		return errors.New(strings.Join(errs, ", "))
	}
}

func (g *ReticGene) UseDefaultAllelicRelations() {
	ar := []snake.TraitRelation{}
	switch g.Trait {
	case snake.Albino:
		fallthrough
	case snake.AlbinoBlond:
		fallthrough
	case snake.AlbinoPurple:
		fallthrough
	case snake.AlbinoWhite:
		ar = append(ar, snake.TraitRelation{Allele: snake.Albino, Expression: snake.Codominant})
		ar = append(ar, snake.TraitRelation{Allele: snake.AlbinoBlond, Expression: snake.Codominant})
		ar = append(ar, snake.TraitRelation{Allele: snake.AlbinoPurple, Expression: snake.Codominant})
		ar = append(ar, snake.TraitRelation{Allele: snake.AlbinoWhite, Expression: snake.Codominant})
	case snake.Anerythirstic:
	case snake.Anthrax:
	case snake.Citrus:
	case snake.GeneticStripe:
	case snake.GoldenChild:
	case snake.Ghost:
	case snake.Jaguar:
	case snake.Hypo:
	case snake.Mocha:
	case snake.Motley:
	case snake.Pied:
	case snake.Phantom:
	case snake.Platinum:
	case snake.Sunfire:
	case snake.Tiger:
	case snake.Titanium:
	case snake.WSexGene:
		ar = append(ar, snake.TraitRelation{Allele: snake.ZSexGene, Expression: snake.CompleteDominance})
	case snake.ZSexGene:
		ar = append(ar, snake.TraitRelation{Allele: snake.WSexGene, Expression: snake.Recessive})
		ar = append(ar, snake.TraitRelation{Allele: snake.ZSexGene, Expression: snake.Codominant})
	}
	g.TraitRelations = ar
}
