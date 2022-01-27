package snake

import "testing"

func TestGender(t *testing.T) {
	if Undefined.String() != "undefined" {
		t.Error("expected Undefined to be expressed as undefined")
	}
	if Male.String() != "male" {
		t.Error("expected Male to be expressed as male")
	}
	if Female.String() != "female" {
		t.Error("expected Female to be expressed as female")
	}
	if Intersex.String() != "intersex" {
		t.Error("expected Intersex to be expressed as intersex")
	}
	if Unsexed.String() != "unsexed" {
		t.Error("expected Unsexed to be expressed as unsexed")
	}
	if Undefined.String() != "undefined" {
		t.Error("expected Undefined to be expressed as undefined")
	}
}
