package retic

import (
	"testing"
)

func TestRetic(t *testing.T) {
	r1 := Retic{
		Traits: []Trait{
			{
				Name:        "Sex",
				Description: "Female is ZW and male is ZZ",
				Expression:  SexLinked,
				AllelePair: AllelePair{
					A1: Dominant,
					A2: Recessive,
				},
			},
			{
				Name:        "Golden Child",
				Description: "Pretty",
				Expression:  SexLinked,
				AllelePair: AllelePair{
					A1: Dominant,
					A2: Recessive,
				},
			},
		},
	}
	r2 := Retic{
		Traits: []Trait{
			{
				Name:        "Sex",
				Description: "Female is ZW and male is ZZ",
				Expression:  CompleteDominance,
				AllelePair: AllelePair{
					A1: Recessive,
					A2: Recessive,
				},
			},
			{
				Name:        "Golden Child",
				Description: "Pretty",
				Expression:  CompleteDominance,
				AllelePair: AllelePair{
					A1: Dominant,
					A2: Recessive,
				},
			},
		},
	}

	res, err := PredictPairing(r1, r2)
	if err != nil {
		t.Error("expected no error, got error")
	}

	t.Error(res)

}
