package retic

import (
	"github.com/geekmdtravis/reticulous/v1/snake"
)

func NewRetic(traits []snake.Trait) snake.Snake {
	s := snake.Snake{}
	s.Genus = "Malayopython"
	s.Species = "reticulatus"
	s.Subspecies, s.Locales = processTraits(traits)
	for _, trait := range traits {
		s.Genes = append(s.Genes, snake.NewGene(trait, setRelations))
	}
	return s
}

func processTraits(t []snake.Trait) (string, map[snake.Trait]float32) {

	locales := make(map[snake.Trait]uint8)
	cnt := 0
	for _, trait := range t {
		switch trait {
		case snake.LocaleMainland:
			fallthrough
		case snake.LocaleKalatoa:
			fallthrough
		case snake.LocaleMadu:
			fallthrough
		case snake.LocaleKarompa:
			fallthrough
		case snake.LocaleKayuadi:
			fallthrough
		case snake.LocaleTanahjampea:
			fallthrough
		case snake.LocaleSelayer:
			fallthrough
		case snake.LocaleUnknown:
			locales[trait]++
			cnt++
		}
	}

	var subspecies string
	percents := make(map[snake.Trait]float32)
	percents[snake.LocaleMainland] = float32(locales[snake.LocaleMainland]) / float32(cnt) * 100
	percents[snake.LocaleKalatoa] = float32(locales[snake.LocaleKalatoa]) / float32(cnt) * 100
	percents[snake.LocaleMadu] = float32(locales[snake.LocaleMadu]) / float32(cnt) * 100
	percents[snake.LocaleKarompa] = float32(locales[snake.LocaleKarompa]) / float32(cnt) * 100
	percents[snake.LocaleKayuadi] = float32(locales[snake.LocaleKayuadi]) / float32(cnt) * 100
	percents[snake.LocaleTanahjampea] = float32(locales[snake.LocaleTanahjampea]) / float32(cnt) * 100
	percents[snake.LocaleSelayer] = float32(locales[snake.LocaleSelayer]) / float32(cnt) * 100
	percents[snake.LocaleUnknown] = float32(locales[snake.LocaleUnknown]) / float32(cnt) * 100

	if int(percents[snake.LocaleMainland]) == 100 ||
		int(percents[snake.LocaleKalatoa]) == 100 ||
		int(percents[snake.LocaleMadu]) == 100 ||
		int(percents[snake.LocaleKarompa]) == 100 {
		subspecies = "reticulatus"
	} else if int(percents[snake.LocaleKayuadi]) == 100 ||
		int(percents[snake.LocaleTanahjampea]) == 100 {
		subspecies = "jampeanus"
	} else if int(percents[snake.LocaleSelayer]) == 100 {
		subspecies = "saputrai"
	} else if int(percents[snake.LocaleUnknown]) == 100 {
		subspecies = "unknown"
	} else {
		subspecies = "mixed"
	}

	return subspecies, percents
}

func setRelations(t snake.Trait) []snake.TraitRelation {
	ar := []snake.TraitRelation{}
	switch t {
	case snake.Albino:
		fallthrough
	case snake.AlbinoBlond:
		fallthrough
	case snake.AlbinoPurple:
		fallthrough
	case snake.AlbinoWhite:
		ar = append(ar, snake.TraitRelation{Trait: snake.Albino, Expression: snake.Codominant})
		ar = append(ar, snake.TraitRelation{Trait: snake.AlbinoBlond, Expression: snake.Codominant})
		ar = append(ar, snake.TraitRelation{Trait: snake.AlbinoPurple, Expression: snake.Codominant})
		ar = append(ar, snake.TraitRelation{Trait: snake.AlbinoWhite, Expression: snake.Codominant})
	case snake.Anerythirstic:
	case snake.Anthrax:
	case snake.Citrus:
	case snake.GeneticStripe:
	case snake.GoldenChild:
	case snake.Ghost:
	case snake.Jaguar:
	case snake.Hypo:
	case snake.Mocha:
	case snake.Motley:
	case snake.Pied:
	case snake.Phantom:
	case snake.Platinum:
	case snake.Sunfire:
	case snake.Tiger:
	case snake.Titanium:
	case snake.WSexGene:
		ar = append(ar, snake.TraitRelation{Trait: snake.ZSexGene, Expression: snake.CompleteDominance})
	case snake.ZSexGene:
		ar = append(ar, snake.TraitRelation{Trait: snake.WSexGene, Expression: snake.Recessive})
		ar = append(ar, snake.TraitRelation{Trait: snake.ZSexGene, Expression: snake.Codominant})
	}
	return ar
}
