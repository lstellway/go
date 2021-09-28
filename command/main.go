/*
	Package command is used to display verbose, easy-to-understand help pages for your CLI commands.

	The package extends upon the flag Standard library package: https://pkg.go.dev/flag
*/
package command

import (
	"flag"
	"fmt"
	"strings"
)

// A CommandExample represents an example usage of the command.
// You can add examples to a command with the Command.AddExample()
// function to help users learn how to use your command.
type CommandExample struct {
	Description string // Description of the example use-case
	Example     string // Verbatim command that can be run
}

// A CommandSection represents a group of related flags.
// You can add and configure a section with the Command.AddSection()
// function.
type CommandSection struct {
	Flags   []string      // Flag names that are related to this section
	FlagSet *flag.FlagSet // Reference to the command FlagSet
	Name    string        // Name of the section
}

type CommandSubcommand struct {
	Description string // Command description
	Name        string // Command name
}

// A Command is a wrapper around the "flag" Standard library package: https://pkg.go.dev/flag
// This package adds utilities to create more verbose usage screens.
type Command struct {
	Usage func()

	Description string              // Command description
	Examples    []CommandExample    // List of examples
	FlagSet     *flag.FlagSet       // Reference to command FlagSet
	Name        string              // Command name
	Arguments   []string            // Any command-line arguments that are available for the command
	Sections    []CommandSection    // List of sections
	Subcommands []CommandSubcommand // List of subcommands
}

// StringVar adds a flag with a boolean value to the FlagSet and groups within a command seciton
func (s *CommandSection) BoolVar(p *bool, name string, value bool, usage string) {
	s.Flags = append(s.Flags, name)
	s.FlagSet.BoolVar(p, name, value, usage)
}

// StringVar adds a flag with an integer value to the FlagSet and groups within a command seciton
func (s *CommandSection) IntVar(p *int, name string, value int, usage string) {
	s.Flags = append(s.Flags, name)
	s.FlagSet.IntVar(p, name, value, usage)
}

// StringVar adds a flag with a string value to the FlagSet and groups within a command seciton
func (s *CommandSection) StringVar(p *string, name string, value string, usage string) {
	s.Flags = append(s.Flags, name)
	s.FlagSet.StringVar(p, name, value, usage)
}

// Add argument
func (h *Command) AddArgument(name string) {
	h.Arguments = append(h.Arguments, name)
}

// AddSection groups related flags together.
// Grouped flags can help with readability on a usage screen
// when a command has a large number of options.
func (h *Command) AddSection(name string, configure func(s *CommandSection)) {
	section := CommandSection{
		Name:    name,
		FlagSet: h.FlagSet,
	}

	// Callback to configure section
	configure(&section)
	h.Sections = append(h.Sections, section)
}

// AddSubcommand defines available subcommands of the current command.
func (h *Command) AddSubcommand(name string, description string) {
	child := CommandSubcommand{
		Name:        name,
		Description: description,
	}
	h.Subcommands = append(h.Subcommands, child)
}

// AddExample defines example usage of the command.
// Use it to display use cases in the help screen.
func (h *Command) AddExample(description string, example string) {
	e := CommandExample{
		Description: description,
		Example:     example,
	}
	h.Examples = append(h.Examples, e)
}

// defaultUsage defines the default usage screen that
// is displayed when using a "help" flag (-h, -help)
func (h *Command) defaultUsage() {
	var help strings.Builder

	// Description
	if len(h.Description) > 0 {
		fmt.Fprintf(&help, "\n%s\n", h.Description)
	}

	// Command name
	if len(h.FlagSet.Name()) > 0 {
		var a strings.Builder
		for _, name := range h.Arguments {
			fmt.Fprintf(&a, " [%s]", name)
		}

		fmt.Fprintf(&help, "\n\nUsage:\n")

		if len(h.Sections) > 0 {
			fmt.Fprintf(&help, "\n    %s [OPTIONS]%s", h.FlagSet.Name(), a.String())
		} else if len(h.Subcommands) > 0 {
			fmt.Fprintf(&help, "\n    %s COMMAND%s", h.FlagSet.Name(), a.String())
		} else {
			fmt.Fprintf(&help, "\n    %s%s", h.FlagSet.Name(), a.String())
		}

		fmt.Fprintf(&help, "\n")
	}

	// Examples
	if len(h.Examples) > 0 {
		fmt.Fprintf(&help, "\n\nExamples:\n")

		for _, e := range h.Examples {
			if e.Description != "" {
				fmt.Fprintf(&help, "\n    %s:", e.Description)
			}
			fmt.Fprintf(&help, "\n      %s %s\n", h.FlagSet.Name(), e.Example)
		}
	}

	// Sections
	if len(h.Sections) > 0 {
		for _, section := range h.Sections {
			fmt.Fprintf(&help, "\n\n%s:\n", section.Name)

			for _, name := range section.Flags {
				// Get flag
				f := h.FlagSet.Lookup(name)

				// Default value
				defValue := ""
				if f.DefValue != "" {
					defValue = fmt.Sprintf("; Default: %v", f.DefValue)
				}

				// Name, type, default
				fmt.Fprintf(&help, "\n   -%s", f.Name)
				fmt.Fprintf(&help, "\n    Type: %T%s\n", f.DefValue, defValue)

				// Usage
				_, flagUsage := flag.UnquoteUsage(f)
				flagUsage = strings.ReplaceAll(flagUsage, "\n", "\n    ")
				fmt.Fprintf(&help, "    %s\n", flagUsage)
			}
		}
	}

	// Subcommands
	if len(h.Subcommands) > 0 {
		fmt.Fprintf(&help, "\n\nCommands:\n")

		for _, s := range h.Subcommands {
			fmt.Fprintf(&help, "\n    %s", s.Name)
			fmt.Fprintf(&help, "\n      %s\n", s.Description)
		}

		fmt.Fprintf(&help, "\n\nRun '%s COMMAND' for more information on a command.\n", h.FlagSet.Name())
	}

	fmt.Fprintf(h.FlagSet.Output(), help.String())
}

// NewCommand initializes a new Command object.
// The Command can be configured via a configuration callback function.
func NewCommand(name string, description string, configure func(h *Command), flags ...string) (Command, []string) {
	h := Command{
		Name:        name,
		Description: description,
		FlagSet:     flag.NewFlagSet(name, flag.ExitOnError),
	}

	// Wire default usage and link to FlagSet
	h.Usage = h.defaultUsage
	h.FlagSet.Usage = func() {
		h.Usage()
	}
	configure(&h)

	h.FlagSet.Parse(flags)
	return h, h.FlagSet.Args()
}
