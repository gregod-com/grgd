package interfaces

import "github.com/urfave/cli/v2"

// IContext ...
type IContext interface {
	// NumFlags returns the number of flags set
	NumFlags() int
	// Set sets a context flag to a value.
	Set(name, value string) error
	// IsSet determines if the flag was actually set
	IsSet(name string) bool
	// LocalFlagNames returns a slice of flag names used in this context.
	LocalFlagNames() []string
	// FlagNames returns a slice of flag names used by the this context and all of
	// its parent contexts.
	FlagNames() []string
	// Lineage returns *this* context and all of its ancestor contexts in order from
	// child to parent
	Lineage() []*IContext
	// Value returns the value of the flag corresponding to `name`
	Value(name string) interface{}
	// Args returns the command line arguments associated with the context.
	Args() cli.Args
	// NArg returns the number of the command line arguments.
	NArg() int
}
