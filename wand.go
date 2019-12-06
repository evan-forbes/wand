package wand

import (
	"fmt"
)

type Spell interface {
	Cast(Context)
}

// Run searches for and runs the first command in the arguments
// provided at runtime
func Run(ctx Context, spells map[string]Spell, args []string) {
	ctx.ParseArgs(args)
	var foundSpell bool
	for _, input := range args {
		spell, has := spells[input]
		if has {
			spell.Cast(ctx)
			foundSpell = true
			break
		}
	}
	if !foundSpell {
		fmt.Println("No viable commands found")
	}
}

// RunMany searches for and runs any commands in the arguments
// provided at runtime in their own goroutines
func RunMany(ctx Context, spells map[string]Spell, args []string) {
	ctx.ParseArgs(args)
	var foundSpell bool
	for _, input := range args {
		spell, has := spells[input]
		if has {
			foundSpell = true
			go spell.Cast(ctx)
		}
	}
	if !foundSpell {
		fmt.Println("No viable commands found")
	}
}
