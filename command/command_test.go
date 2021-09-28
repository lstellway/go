package command_test

import (
	"os"

	"github.com/lstellway/go/command"
)

var (
	isVerbose, log bool
	logFile        string

	cmd  command.Command
	args []string
)

func ExampleCommand() {
	// Define a command with a list of subcommands
	cmd, args = command.NewCommand("example", "Description of the example command", func(c *command.Command) {
		c.AddSubcommand("list", "List available resources")
		c.AddSubcommand("cp", "Copy a resource to a specified location")
		c.AddSubcommand("rm", "Remove a resource")
	}, os.Args[1:]...)
}

func ExampleCommandExample() {
	// Define a command with a list of examples
	cmd, args = command.NewCommand("example", "Description of the example command", func(c *command.Command) {
		c.AddExample("List available resources", "example list")
		c.AddExample("Copy a resource to a specified location", "example cp")
		c.AddExample("Remove a resource", "example rm")
	}, os.Args[1:]...)
}

func ExampleCommandSection() {
	// Define a command with sections of options
	cmd, args = command.NewCommand("example", "Description of the example command", func(c *command.Command) {
		// General options
		c.AddSection("General Options", func(s *command.CommandSection) {
			s.BoolVar(&isVerbose, "verbose", false, "Show verbose output")
			s.BoolVar(&log, "log", true, "Enable logging")
		})

		// Advanced options
		c.AddSection("Advanced Options", func(s *command.CommandSection) {
			s.StringVar(&logFile, "logger", "/dev/stdout", "File to log output")
		})
	}, os.Args[1:]...)
}

func ExampleCommandSubcommand() {
	// Define a command with a list of subcommands
	cmd, args = command.NewCommand("example", "Description of the example command", func(c *command.Command) {
		c.AddSubcommand("list", "List available resources")
		c.AddSubcommand("cp", "Copy a resource to a specified location")
		c.AddSubcommand("rm", "Remove a resource")
	}, os.Args[1:]...)
}
