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

func (a Allele) String() string {
	switch a {
	case Albino:
		return "albino"
	case AlbinoBlond:
		return "albino blond"
	case AlbinoPurple:
		return "albino purple"
	case AlbinoWhite:
		return "albino white"
	case Anerythirstic:
		return "anerythirstic"
	case Anthrax:
		return "anthrax"
	case Citrus:
		return "citrus"
	case GeneticStripe:
		return "genetic stripe"
	case GoldenChild:
		return "golden child"
	case Ghost:
		return "ghost"
	case Jaguar:
		return "jaguar"
	case Hypo:
		return "hypo"
	case Mocha:
		return "mocha"
	case Pied:
		return "pied"
	case Phantom:
		return "phantom"
	case Platinum:
		return "platinum"
	case Sunfire:
		return "sunfire"
	case Tiger:
		return "tiger"
	case Titanium:
		return "titanium"
	case ZW:
		return "female (ZW)"
	case ZZ:
		return "male (ZZ)"
	}
	return "unknown"
}

type AllelicRelation struct {
	Allele    Allele
	Recessive bool
	Dominance Dominance
}
