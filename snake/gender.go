package snake

type Gender uint8

const (
	Undefined Gender = iota
	Male
	Female
	Intersex
	Unsexed
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	case Intersex:
		return "intersex"
	case Unsexed:
		return "unsexed"
	}
	return "undefined"
}
