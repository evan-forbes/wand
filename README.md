# wand

THE simplest, most "anti-framework" command line interface making tool for go. Wand iterates through some slice of strings, and if one of those strings is assigned to a spell (a sub command), then it will run that spell's Cast method.

## Usage

```golang
type FirstCommand struct{}

// Cast is like a mini main func for each of your sub-commands.
func (ms *FirstCommand) Cast(wand.Context) {
    // R.I.P. here lies your code
    fmt.Println("Hello, friend!")
}

func main() {
    ctx := wand.NewDefaultContext(context.Background())
    spells := map[string]wand.Spell{
        "FirstCommand": &FirstCommand{},
    }
    wand.Run(ctx, spells, os.Args[1:])
}
```
```
$ go run main.go FirstCommand
// Hello, friend!
```
Any additional arguments or flags are stored in the wand.Context, and accessable by using the methods Args() and Flags(), both of witch (ha!) return a `map[string]string`. If you decide to use the wand.DefaultContext, then anything following the flag or argument is stored as the value, with the argument or flag stored as the key. ie.
```
go run main.go command example/path/to/thing -shortFlag --longFlag=Thing
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
If that parsing mechanism doesn't work for you, wand.Context is an interface, and I'd love to put another in the library. **(you should forreal make another one and submit a PR)**

I would adore any contributions or ideas, in any form, including submitting an issue.

Made with witchly love -
    Evan

PS: Remember, to some, a wand is just a stick.