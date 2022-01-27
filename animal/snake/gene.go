package snake

type Gene struct {
	Allele       Allele
	Relationship []AllelicRelation
	Inheritance  Inheritance
	Zygosity     Zygosity
	Description  string
	Percentage   float64
}

type Dominance uint8

const (
	Complete Dominance = iota
	Incomplete
	Codominant
)

func (d Dominance) String() string {
	switch d {
	case Complete:
		return "complete"
	case Incomplete:
		return "incomplete"
	case Codominant:
		return "codominant"
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
