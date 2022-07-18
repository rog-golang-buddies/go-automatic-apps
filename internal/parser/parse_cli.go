package parser

import (
	"fmt"
	"strings"
)

func ParseCliArguments(arguments []string) ([]string, map[string]interface{}, error) {
	var commands []string

	flags := map[string]interface{}{}

	flagsStarted := false
	var lastValueWasFlag bool
	var lastValueFlag string

	for _, value := range arguments {
		if !flagsStarted && !strings.Contains(value, "-") {
			commands = append(commands, value)
			continue
		}

		// Check if the string is a flag
		var isFlag bool

		if strings.Contains(value, "-") {
			flagsStarted = true
			isFlag = true
		}

		// add value to the last value flag
		if lastValueWasFlag {
			// The current value is also a flag so we just add true to the flags value
			flag := strings.ReplaceAll(lastValueFlag, "-", "")

			if flag == "" {
				return nil, nil, fmt.Errorf("empty flag was passed in")
			}

			if flags[flag] != nil {
				return nil, nil, fmt.Errorf("flag %v was set multiple times", flag)
			}
			if isFlag {
				flags[flag] = true
			} else {
				flags[flag] = value
			}
			lastValueWasFlag = false
			lastValueFlag = ""

		} else {
			// the last value was not a flag
			// if this value is also not a flag the user entered in 2 values for a single flag
			if !isFlag {
				return nil, nil, fmt.Errorf("invalid number of parameters. Unexpected value %v", value)
			}
		}

		// Used for next iteration
		if isFlag {
			lastValueFlag = value
			lastValueWasFlag = true
		}
	}

	if lastValueWasFlag {
		flag := strings.ReplaceAll(lastValueFlag, "-", "")
		if flag == "" {
			return nil, nil, fmt.Errorf("empty flag was passed in")
		}
		if flags[flag] != nil {
			return nil, nil, fmt.Errorf("flag %v was set multiple times", flag)
		}
		flags[flag] = true
	}

	if len(commands) == 0 {
		return nil, nil, fmt.Errorf("no action was provided")
	}

	return commands, flags, nil
}
