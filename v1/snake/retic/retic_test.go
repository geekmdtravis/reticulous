package retic

import (
	"testing"

	"github.com/geekmdtravis/reticulous/snake"
)

func TestRetic(t *testing.T) {
	mainlandTraits := []snake.Trait{
		snake.LocaleMainland,
		snake.LocaleMainland,
		snake.Albino,
	}
	s := NewRetic(mainlandTraits)

	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "reticulatus" {
		t.Error("expected reticulatus, got", s.Subspecies)
	}

	kalatoaTraits := []snake.Trait{
		snake.LocaleKalatoa,
		snake.AlbinoPurple,
	}
	s = NewRetic((kalatoaTraits))
	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "reticulatus" {
		t.Error("expected reticulatus, got", s.Subspecies)
	}

	maduTraits := []snake.Trait{
		snake.LocaleMadu,
		snake.LocaleMadu,
		snake.LocaleMadu,
		snake.LocaleMadu,
		snake.AlbinoPurple,
	}
	s = NewRetic((maduTraits))
	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "reticulatus" {
		t.Error("expected reticulatus, got", s.Subspecies)
	}
	if s.Locales[snake.LocaleMadu] != 100 {
		t.Error("expected 100, got", s.Locales[snake.LocaleMadu])
	}

	karompaTraits := []snake.Trait{
		snake.LocaleKarompa,
		snake.AlbinoPurple,
	}
	s = NewRetic((karompaTraits))
	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "reticulatus" {
		t.Error("expected reticulatus, got", s.Subspecies)
	}
	if s.Locales[snake.LocaleKarompa] != 100 {
		t.Error("expected 100, got", s.Locales[snake.LocaleKarompa])
	}

	kayuadiTraits := []snake.Trait{
		snake.LocaleKayuadi,
		snake.AlbinoPurple,
	}
	s = NewRetic((kayuadiTraits))
	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "jampeanus" {
		t.Error("expected jampeanus, got", s.Subspecies)
	}
	if s.Locales[snake.LocaleKayuadi] != 100 {
		t.Error("expected 100, got", s.Locales[snake.LocaleKayuadi])
	}

	selayerTraits := []snake.Trait{
		snake.LocaleSelayer,
		snake.AlbinoPurple,
	}
	s = NewRetic((selayerTraits))
	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "saputrai" {
		t.Error("expected saputrai, got", s.Subspecies)
	}
	if s.Locales[snake.LocaleSelayer] != 100 {
		t.Error("expected 100, got", s.Locales[snake.LocaleSelayer])
	}

	unkTraits := []snake.Trait{
		snake.LocaleUnknown,
		snake.AlbinoPurple,
	}
	s = NewRetic((unkTraits))
	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "unknown" {
		t.Error("expected unknown, got", s.Subspecies)
	}
	if s.Locales[snake.LocaleUnknown] != 100 {
		t.Error("expected 100, got", s.Locales[snake.LocaleUnknown])
	}

	mixedTraits := []snake.Trait{
		snake.LocaleKalatoa,
		snake.LocaleMainland,
		snake.LocaleMadu,
		snake.AlbinoPurple,
	}
	s = NewRetic((mixedTraits))
	if s.Genus != "Malayopython" {
		t.Error("expected Malayopython, got", s.Genus)
	}
	if s.Species != "reticulatus" {
		t.Error("expected reticulatus, got", s.Species)
	}
	if s.Subspecies != "mixed" {
		t.Error("expected mixed, got", s.Subspecies)
	}
	if int(s.Locales[snake.LocaleKalatoa]) != 33 {
		t.Error("expected 33, got", int(s.Locales[snake.LocaleKalatoa]))
	}
	if int(s.Locales[snake.LocaleMainland]) != 33 {
		t.Error("expected 33, got", int(s.Locales[snake.LocaleMainland]))
	}
	if int(s.Locales[snake.LocaleMadu]) != 33 {
		t.Error("expected 33, got", int(s.Locales[snake.LocaleMadu]))
	}
}
