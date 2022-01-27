package snake

import (
	"testing"
	"time"

	"github.com/geekmdtravis/reticulous/v1/blockchain"
)

func makeValidSnake() Snake {
	locales := make(map[Trait]float32)
	locales[LocaleKayuadi] = 0.5
	locales[LocaleMadu] = 0.5

	genes := make([]Gene, 4)
	g1 := Gene{
		Trait: GoldenChild,
		TraitRelations: []TraitRelation{{
			Trait:      Albino,
			Expression: Recessive,
		}},
	}
	g2 := Gene{
		Trait: Albino,
		TraitRelations: []TraitRelation{{
			Trait:      GoldenChild,
			Expression: IncompleteDominance,
		}},
	}
	genes = append(genes, g1, g2)

	s := Snake{
		DateOfBirth: time.Date(2010, time.February, 13, 0, 0, 0, 0, time.UTC),
		Name:        "test snake",
		NFTAddress:  "0x1234567890123456789012345678901234567890",
		NFTNetwork:  blockchain.Ethereum,
		Gender:      Male,
		Sire:        &Snake{},
		Dam:         &Snake{},
		Genes:       genes,
		Locales:     locales,
		Taxonomy: Taxonomy{
			Genus:      "Malayopython",
			Species:    "reticulatus",
			Subspecies: "reticulatus",
		},
	}
	return s
}
func TestSnake(t *testing.T) {
	s := makeValidSnake()
	err := s.Validate()
	if err != nil {
		t.Error("expected nil with valid snake, got", err)
	}

	s.DateOfBirth = time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	err = s.Validate()
	if err == nil {
		t.Error("expected error with invalid DateOfBirth, got nil")
	}

	s = makeValidSnake()
	s.Name = ""
	err = s.Validate()
	if err == nil {
		t.Error("expected error with invalid Name, got nil")
	}

	s = makeValidSnake()
	s.NFTAddress = ""
	err = s.Validate()
	if err == nil {
		t.Error("expected error with invalid NFTAddress, got nil")
	}

	s = makeValidSnake()
	s.NFTNetwork = blockchain.Undefined
	err = s.Validate()
	if err == nil {
		t.Error("expected error with invalid NFTNetwork, got nil")
	}

	s = makeValidSnake()
	s.Gender = Undefined
	err = s.Validate()
	if err == nil {
		t.Error("expected error with invalid Gender, got nil")
	}

	s = makeValidSnake()
	s.Genes = []Gene{}
	err = s.Validate()
	if err == nil {
		t.Error("expected error with empty Gene set, got nil")
	}

	s = makeValidSnake()
	s.Locales = make(map[Trait]float32)
	err = s.Validate()
	if err == nil {
		t.Error("expected error with empty Locales, got nil")
	}

	s = makeValidSnake()
	s.Genus = ""
	err = s.Validate()
	if err == nil {
		t.Error("expected error with empty Genus, got nil")
	}

	s = makeValidSnake()
	s.Species = ""
	err = s.Validate()
	if err == nil {
		t.Error("expected error with empty Species, got nil")
	}

	s = makeValidSnake()
	s.Subspecies = ""
	err = s.Validate()
	if err == nil {
		t.Error("expected error with empty Subspecies, got nil")
	}

}
