package wand

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

type testSpell struct {
	SpellCasted bool
}

func (ts *testSpell) Cast(ctx Context) {
	fmt.Println("Casting spell")
	ts.SpellCasted = true
}

func newTestSpell() *testSpell {
	return &testSpell{}
}

func TestRun(t *testing.T) {
	ctx := NewDefaultContext(context.Background())
	tspell := newTestSpell()
	spells := map[string]Spell{
		"test": tspell,
	}
	Run(ctx, spells, []string{"placeholder", "test"})
	if !tspell.SpellCasted {
		t.Error(errors.New("spell not casted"))
	}
}
