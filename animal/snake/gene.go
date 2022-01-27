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

type Zygosity uint8

const (
	Homozygous Zygosity = iota
	Heterozygous
	Hemizygous
	Nullizygous
	NotApplicable
)

func (z Zygosity) String() string {
	switch z {
	case Homozygous:
		return "homozygous"
	case Heterozygous:
		return "heterozygous"
	case Hemizygous:
		return "hemizygous"
	case Nullizygous:
		return "nullizygous"
	case NotApplicable:
		return "not applicable"
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

type Gene struct {
	Trait          Trait
	traitRelations []TraitRelation
	Inheritance    Inheritance
	Zygosity       Zygosity
	Description    string
	Percentage     float64
}

func (g Gene) Validate() error {
	errs := make([]string, 0, 3)

	if len(g.traitRelations) == 0 {
		errs = append(errs, "no allelic expression relationships are defined")
	}

	if g.Inheritance == Mendelian && !(int(g.Percentage) == 50) {
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

func (g *Gene) UseDefaultAllelicRelations() {
	ar := []TraitRelation{}
	switch g.Trait {
	case Albino:
		fallthrough
	case AlbinoBlond:
		fallthrough
	case AlbinoPurple:
		fallthrough
	case AlbinoWhite:
		ar = append(ar, TraitRelation{Allele: Albino, Expression: Codominant})
		ar = append(ar, TraitRelation{Allele: AlbinoBlond, Expression: Codominant})
		ar = append(ar, TraitRelation{Allele: AlbinoPurple, Expression: Codominant})
		ar = append(ar, TraitRelation{Allele: AlbinoWhite, Expression: Codominant})
	case Anerythirstic:
	case Anthrax:
	case Citrus:
	case GeneticStripe:
	case GoldenChild:
	case Ghost:
	case Jaguar:
	case Hypo:
	case Mocha:
	case Motley:
	case Pied:
	case Phantom:
	case Platinum:
	case Sunfire:
	case Tiger:
	case Titanium:
	case WSexGene:
		ar = append(ar, TraitRelation{Allele: ZSexGene, Expression: CompleteDominance})
	case ZSexGene:
		ar = append(ar, TraitRelation{Allele: WSexGene, Expression: Recessive})
		ar = append(ar, TraitRelation{Allele: ZSexGene, Expression: Codominant})
	}
	g.traitRelations = ar
}
