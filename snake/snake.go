package snake

import (
	"errors"
	"strings"
	"time"

	"github.com/geekmdtravis/reticulous/blockchain"
)

type Snake struct {
	DateOfBirth time.Time
	Name        string
	NFTAddress  string
	NFTNetwork  blockchain.Network
	Gender      Gender
	Sire        *Snake
	Dam         *Snake
	Genes       []Gene
	Locales     map[Trait]float32
	Taxonomy
}

func (s Snake) Validate() error {
	errs := make([]string, 0, 10)
	if s.DateOfBirth.IsZero() {
		errs = append(errs, "DateOfBirth is required")
	}
	if s.Name == "" {
		errs = append(errs, "Name is required")
	}
	if s.NFTAddress == "" {
		errs = append(errs, "NFTAddress is required")
	}
	if s.NFTNetwork == blockchain.Undefined {
		errs = append(errs, "NFTNetwork is required")
	}
	if s.Gender == Undefined {
		errs = append(errs, "Gender is required")
	}
	if len(s.Genes) == 0 {
		errs = append(errs, "Genes need to be populated")
	}
	if len(s.Locales) == 0 {
		errs = append(errs, "Locales need to be populated")
	}
	if s.Genus == "" {
		errs = append(errs, "Genus shoudl be populated from traits")
	}
	if s.Species == "" {
		errs = append(errs, "Species should be populated from traits")
	}
	if s.Subspecies == "" {
		errs = append(errs, "Subspecies should be populated from traits")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ","))
	}
	return nil

}
