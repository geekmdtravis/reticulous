package snake

import (
	"errors"
	"strings"
)

type Expression uint8

const (
	CompleteDominance Expression = iota
	IncompleteDominance
	Codominant
	Recessive
)

func (d Expression) String() string {
	switch d {
	case CompleteDominance:
		return "complete dominance"
	case IncompleteDominance:
		return "incomplete dominance"
	case Codominant:
		return "codominant"
	case Recessive:
		return "recessive"
	}
	return "unknown"
}

type Inheritance uint8

const (
	Mendelian Inheritance = iota
	NonMendelian
	SexLinkedZZMale
	SexLinkedZWFemale
)

func (i Inheritance) String() string {
	switch i {
	case Mendelian:
		return "mendelian"
	case NonMendelian:
		return "non-mendelian"
	case SexLinkedZZMale:
		return "male sex linked (ZZ)"
	case SexLinkedZWFemale:
		return "female sex linked (ZW)"
	}
	return "unknown"
}

// Plan to handle percentage by adding multiple copies of a non-mendelian gene
// E.g. [Jampea, Jampea, Jampea] and [Mainland, Mainland] would be 60% Jampea
type Gene struct {
	Trait          Trait
	TraitRelations []TraitRelation
	Inheritance    Inheritance
	Description    string
	// Percentage     float64
}

func (g Gene) Validate() error {
	errs := make([]string, 0, 3)

	if len(g.TraitRelations) == 0 {
		errs = append(errs, "no allelic expression relationships are defined")
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
