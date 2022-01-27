package snake

import (
	"errors"
	"strings"
)

type Zygosity uint8

const (
	UndeclaredZygosity Zygosity = iota
	Heterozygous
	Homozygous
	Hemizygous
	NotApplicable
)

func (z Zygosity) String() string {
	switch z {
	case Heterozygous:
		return "heterozygous"
	case Homozygous:
		return "homozygous"
	case Hemizygous:
		return "hemizygous"
	case NotApplicable:
		return "not applicable"
	default:
		return "undeclared"
	}
}

type Expression uint8

const (
	UndeclaredExpression Expression = iota
	CompleteDominance
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
	return "undeclared"
}

type Inheritance uint8

const (
	UndeclaredInheritance Inheritance = iota
	Mendelian
	NonMendelian
	SexLinked
)

func (i Inheritance) String() string {
	switch i {
	case Mendelian:
		return "mendelian"
	case NonMendelian:
		return "non-mendelian"
	case SexLinked:
		return "female sex linked"
	}
	return "undeclared"
}

type Phenotype uint8

const (
	UndeclaredPhenotype Phenotype = iota
	Size
	Color
	Pattern
)

func (p Phenotype) String() string {
	switch p {
	case Size:
		return "size"
	case Color:
		return "color"
	case Pattern:
		return "pattern"
	default:
		return "undeclared"
	}
}

// TODO: Plan to handle percentage by adding multiple copies of a non-mendelian gene. E.g. [Jampea, Jampea, Jampea] and [Mainland, Mainland] would be 60% Jampea
type Gene struct {
	Phenotype      Phenotype
	Trait          Trait
	TraitRelations []TraitRelation
}

func NewGene(t Trait, cb func(t Trait) []TraitRelation) Gene {
	return Gene{
		Trait:          t,
		TraitRelations: cb(t),
	}
}

func (g Gene) Validate() error {
	errs := make([]string, 0, 3)

	if g.Phenotype == UndeclaredPhenotype {
		errs = append(errs, "phenotype is undeclared")
	}

	if g.Trait == UndeclaredTrait {
		errs = append(errs, "trait is undeclared")
	}

	if len(g.TraitRelations) == 0 {
		errs = append(errs, "no allelic expression relationships are defined")
	}

	if len(errs) == 0 {
		return nil
	} else {
		return errors.New(strings.Join(errs, ", "))
	}
}

func (g Gene) Inheritance() Inheritance {
	switch g.Trait {
	case Albino:
		fallthrough
	case AlbinoBlond:
		fallthrough
	case AlbinoPurple:
		fallthrough
	case AlbinoWhite:
		fallthrough
	case Anerythirstic:
		fallthrough
	case Anthrax:
		fallthrough
	case Citrus:
		fallthrough
	case GeneticStripe:
		fallthrough
	case GoldenChild:
		fallthrough
	case Ghost:
		fallthrough
	case Jaguar:
		fallthrough
	case Hypo:
		fallthrough
	case Mocha:
		fallthrough
	case Motley:
		fallthrough
	case Pied:
		fallthrough
	case Phantom:
		fallthrough
	case Platinum:
		fallthrough
	case Sunfire:
		fallthrough
	case Tiger:
		fallthrough
	case Titanium:
		return Mendelian
	case WSexGene:
		fallthrough
	case ZSexGene:
		return SexLinked
	case LocaleMainland:
		fallthrough
	case LocaleKalatoa:
		fallthrough
	case LocaleMadu:
		fallthrough
	case LocaleKarompa:
		fallthrough
	case LocaleKayuadi:
		fallthrough
	case LocaleTanahjampea:
		fallthrough
	case LocaleSelayer:
		fallthrough
	case LocaleUnknown:
		fallthrough
	case PosDwarf:
		fallthrough
	case PosSuperDwarf:
		fallthrough
	default:
		return NonMendelian
	}
}
