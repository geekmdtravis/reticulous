package snake

import (
	"errors"
	"strings"
)

type Expression uint8

const (
	Complete Expression = iota
	Incomplete
	Codominant
	Recessive
)

func (d Expression) String() string {
	switch d {
	case Complete:
		return "complete"
	case Incomplete:
		return "incomplete"
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
	Allele           Allele
	AllelicRelations []AllelicRelation
	Inheritance      Inheritance
	Zygosity         Zygosity
	Description      string
	Percentage       float64
}

func (g Gene) Validate() error {
	errs := make([]string, 0, 3)

	if len(g.AllelicRelations) == 0 {
		errs = append(errs, "no allelic expression relationships are defined")
	}

	pctg50 := int(g.Percentage) == 50
	pctg100 := int(g.Percentage) == 100
	if g.Inheritance == Mendelian && !(pctg50 || pctg100) {
		errs = append(errs, "with Mendelian inheritance, percentage must be 50 or 100")
	}

	if g.Percentage <= 0 || g.Percentage > 100 {
		errs = append(errs, "percentage must be greater than 0 and less than or equal to 100")
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
