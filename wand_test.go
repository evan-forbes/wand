package wand

import (
	"context"
	"errors"
	"fmt"
	"strings"
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

type contSpell struct{}

func (cs *contSpell) Cast(Context) {

}

func _parseFlag(flag string) (string, string) {
	if !strings.Contains(flag, "=") {
		return strings.Trim(flag, "-"), ""
	}
	s := strings.Split(flag, "=")
	return strings.Trim(s[0], "-"), s[1]
}

func TestParseFlags(t *testing.T) {
	inputs := []string{"-flag", "--longFlag", "--longFlagWith=Tag", "--"}
	expected := map[string]string{
		"flag":         "",
		"longFlag":     "",
		"longFlagWith": "Tag",
		"":             "",
	}
	for _, input := range inputs {
		key, val := _parseFlag(input)
		t.Log(fmt.Sprintf("results key: %q val: %q", key, val))
		expVal, has := expected[key]
		if !has {
			t.Errorf("flag not found: wanted %s, got %s", expVal, val)
		}
		if expVal != val {
			t.Errorf("unexpected value from flags: should be %s, got %s", expVal, val)
		}
	}
}

func TestContextFlags(t *testing.T) {
	inputs := [][]string{
		[]string{"-flag", "--longFlag", "--longFlagWith=Tag", "--"},
	}
	expected := []map[string]string{
		map[string]string{
			"flag":         "",
			"longFlag":     "",
			"longFlagWith": "Tag",
			"":             "",
		},
	}
	for i, input := range inputs {
		ctx := NewDefaultContext(context.Background())
		ctx.ParseArgs(input)
		flags := ctx.Flags()
		if len(flags) != len(expected[i]) {
			t.Log(flags)
			t.Error("unexpected number of flags")
		}
		for key, val := range expected[i] {
			inVal, has := flags[key]
			if !has {
				t.Errorf("expected value not found in flags: %s %s", val, inVal)
			}
			if inVal != val {
				t.Errorf("unexpected value from flags: should be %s, got %s", "val", "inVal")
			}
		}
	}
}

func TestContextCommands(t *testing.T) {
	inputs := [][]string{
		[]string{"command", "path"},
		[]string{"command", "otherCommand"},
	}
	expected := []map[string]string{
		map[string]string{
			"command": "path",
			"path":    "",
		},
		map[string]string{
			"command":      "otherCommand",
			"otherCommand": "",
		},
		map[string]string{},
	}
	for i, input := range inputs {
		ctx := NewDefaultContext(context.Background())
		ctx.ParseArgs(input)
		args := ctx.Args()
		if len(args) != len(expected[i]) {
			t.Log(args)
			t.Error("unexpected number of arguments")
		}
		for key, val := range expected[i] {
			inVal, has := args[key]
			if !has {
				t.Errorf("expected value not found in args: %s %s", val, inVal)
			}
			if inVal != val {
				t.Errorf("unexpected value from args: should be %s, got %s", "val", "inVal")
			}
		}
	}
}
