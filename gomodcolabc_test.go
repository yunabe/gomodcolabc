package gomodcolabc

import "testing"

func TestVersionOfModuleA(t *testing.T) {
	got := VersionOfModuleA()
	want := "v1.1.2"
	if got != want {
		t.Errorf("VersionOfModuleA = %q; want %q", got, want)
	}
}
