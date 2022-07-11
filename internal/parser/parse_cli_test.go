package parser

import (
	"testing"
)

func TestParseCliArguments(t *testing.T) {

	t.Run("Correctly gets commands", func(t *testing.T) {
		commands, _, _ := ParseCliArguments([]string{"test_command1", "test_command2"})

		if len(commands) != 2 {
			t.Fatalf("Invalid command length. Expected %v got %v\n", 2, len(commands))
		}

		if commands[0] != "test_command1" {
			t.Fatalf("Error in commands expected %v got %v\n", "test_command1", commands[0])
		}

		if commands[1] != "test_command2" {
			t.Fatalf("Error in commands expected %v got %v\n", "test_command2", commands[1])
		}
	})

	t.Run("Correctly creates flags that have values", func(t *testing.T) {
		_, flags, _ := ParseCliArguments([]string{"test_command1", "--foo", "bar", "--baz", "foo"})

		if len(flags) != 2 {
			t.Fatalf("Invalid flags length. Expected %v got %v\n", 2, len(flags))
		}

		if flags["foo"] != "bar" {
			t.Fatalf("Error in flags expected %v got %v\n", "bar", flags["foo"])
		}

		if flags["baz"] != "foo" {
			t.Fatalf("Error in flags expected %v got %v\n", "foo", flags["baz"])
		}
	})

	t.Run("Correctly creates flags that are booleans", func(t *testing.T) {
		_, flags, _ := ParseCliArguments([]string{"test_command1", "--foo", "--baz", "foo"})

		if len(flags) != 2 {
			t.Fatalf("Invalid flags length. Expected %v got %v\n", 2, len(flags))
		}

		if flags["foo"] != true {
			t.Fatalf("Error in flags expected %v got %v\n", true, flags["foo"])
		}

		if flags["baz"] != "foo" {
			t.Fatalf("Error in flags expected %v got %v\n", "foo", flags["baz"])
		}
	})
	t.Run("Correctly errors if you enter too many values for a given flag", func(t *testing.T) {
		_, _, err := ParseCliArguments([]string{"test_command1", "--foo", "value1", "value2", "--bar"})

		if err == nil {
			t.Fatalf("Error was expected but not received")
		}

		if err.Error() != "invalid number of parameters. Unexpected value value2" {
			t.Fatalf("Error message incorrect. Expected \"%v\" got \"%v\"", "invalid number of parameters. Unexpected value value2", err.Error())
		}
	})
	t.Run("Correctly errors if you try to pass in the same flag twice", func(t *testing.T) {
		_, _, err := ParseCliArguments([]string{"test_command1", "--foo", "value1", "--foo", "--bar"})

		if err == nil {
			t.Fatalf("Error was expected but not received")
		}

		if err.Error() != "flag foo was set multiple times" {
			t.Fatalf("Error message incorrect. Expected \"%v\" got \"%v\"", "flag foo was set multiple times", err.Error())
		}
	})
	t.Run("Correctly errors if you try to pass in the same flag twice, duplicate flag is the last value", func(t *testing.T) {
		_, _, err := ParseCliArguments([]string{"test_command1", "--foo", "value1", "--bar", "--foo"})

		if err == nil {
			t.Fatalf("Error was expected but not received")
		}

		if err.Error() != "flag foo was set multiple times" {
			t.Fatalf("Error message incorrect. Expected \"%v\" got \"%v\"", "flag foo was set multiple times", err.Error())
		}
	})
	t.Run("Correctly errors if you pass in empty -", func(t *testing.T) {
		_, _, err := ParseCliArguments([]string{"test_command1", "--", "value1", "--bar", "--foo"})

		if err == nil {
			t.Fatalf("Error was expected but not received")
		}

		if err.Error() != "empty flag was passed in" {
			t.Fatalf("Error message incorrect. Expected \"%v\" got \"%v\"", "empty flag was passed in", err.Error())
		}
	})

}
