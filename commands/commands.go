package commands

// Command represents an executable cli sub/command.
type Command interface {

	// Name of the command
	Name() string

	// Execute the command's logic with the provided arguments and flags.
	Exec([]string) error
}

// Commands returns a list of all available commands.
func Commands() []Command {
	return []Command{
		new(Init),
	}
}
