package snake

type Allele uint16

const (
	Albino Allele = iota
	AlbinoBlond
	AlbinoPurple
	AlbinoWhite
	Anerythirstic
	Anthrax
	Citrus
	GeneticStripe
	GoldenChild
	Ghost
	Jaguar
	Hypo
	Mocha
	Motley
	Pied
	Phantom
	Platinum
	Sunfire
	Tiger
	Titanium
	ZW
	ZZ
)

type Dominance uint8

const (
	Complete Dominance = iota
	Incomplete
	Codominant
)

type Zygosity uint8

const (
	Homozygous Zygosity = iota
	Heterozygous
	Hemizygous
	Nullizygous
	NotApplicable
)

type AllelicRelation struct {
	Allele    Allele
	Recessive bool
	Dominance Dominance
}

type Inheritance uint8

const (
	Mendelian Inheritance = iota
	NonMendelian
	SexLinkedZZMale
	SexLinkedZWFemale
)

type Gene struct {
	Allele       Allele
	Relationship []AllelicRelation
	Inheritance  Inheritance
	Zygosity     Zygosity
	Description  string
	Percentage   float64
}
