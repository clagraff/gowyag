package commands

import "testing"

func TestCommands(t *testing.T) {
	availableCommands := Commands()

	expectedLen := 1
	actualLen := len(availableCommands)

	if actualLen != expectedLen {
		t.Errorf(
			"expected length: %d did not equal actual length: %d",
			expectedLen,
			actualLen,
		)
	}
}
