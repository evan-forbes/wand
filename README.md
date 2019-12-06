# wand

The simplest, most "anti-framework", least amount of code, command line interface making tool for go. `wand.Run` iterates through some slice of `string`, and if one of those strings is assigned to a spell (a sub command), then it will run that spell's `Cast` method. That's it!

Should you use this for a big gnarly app? No, probably not. Just want to quickly call a few different functions without making multiple binaries or using a larger more cumbersome framework? Yes, that's a pretty solid use case, or that's at least what I use it for.

## Usage

```golang
type PrintCommand struct{
    Message string
}

// Cast is like a mini main func for each of your sub-commands.
func (pc *PrintCommand) Cast(ctx wand.Context) {
    // R.I.P. here lies your code. ie
    fmt.Println(pc.Message)
}

func main() {
    ctx := wand.NewDefaultContext(context.Background())
    spells := map[string]wand.Spell{
        "PrintCommand": &PrintCommand{"Hello, Friend"},
    }
    wand.Run(ctx, spells, os.Args[1:])
}
```
```
$ go run main.go PrintCommand
// Hello, friend!
```
Any additional arguments or flags are stored in the `wand.Context`, and accessible to your `Cast` method by using the methods `ctx.Args()` and `ctx.Flags()`, both of witch (ha!) return a `map[string]string`. If you decide to use the `wand.DefaultContext`, then anything following the flag or argument is stored as the value, with the argument or flag stored as the key. ie.
```
$ go run main.go command example/path/to/thing -shortFlag --longFlag=Thing
```
would result in
```golang
flags := ctx.Flags()
flags == map[string]string{
    "shortFlag": "",
    "longFlag": "Thing",
}
// where args would be
args := ctx.Args()
args == map[string]string{
    "command": "example/path/to/thing",
}
```
Checking for flags or other data is as easy as
```golang
func (mc *MyCommand) Cast(ctx wand.Context) {
    flags := ctx.Flags()
    
    _, hasFlag := flags["flag"]
    _, hasOtherFlag := flags["otherFlag"]
    switch {
    case hasFlag:
        // do something special
    case hasOtherFlag:
        // do something extra special
    }
}
```

If that parsing mechanism doesn't work for you, `wand.Context` is an interface, and I'd love to put another in the library. **(you should forreal make another one and submit a PR, I WILL ACCEPT IT AND WE CAN BE *super cool* INTERNET FRIENDS)**

There's still a decent amount of work that needs to be done for anyone to actually want to use this other than myself, and I would adore any contributions or ideas, in any form, including submitting an issue. Sooner or later, I'll try and add bash like auto complete of commands.

Made with witchly love -
    Evan

PS: Remember, to some, a wand is just a stick.