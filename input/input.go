package input

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bitrise-io/go-utils/pathutil"
)

// ValidateIfNotEmpty ...
func ValidateIfNotEmpty(input string) error {
	if input == "" {
		return fmt.Errorf("parameter not specified")
	}
	return nil
}

// ValidateWithOptions ...
func ValidateWithOptions(value string, options ...string) error {
	if err := ValidateIfNotEmpty(value); err != nil {
		return err
	}
	for _, option := range options {
		if option == value {
			return nil
		}
	}
	return fmt.Errorf("invalid parameter: %s, available: %v", value, options)
}

// ValidateIfPathExists ...
func ValidateIfPathExists(input string) error {
	if err := ValidateIfNotEmpty(input); err != nil {
		return err
	}
	if exist, err := pathutil.IsPathExists(input); err != nil {
		return fmt.Errorf("failed to check if path exist at: %s, error: %s", input, err)
	} else if !exist {
		return fmt.Errorf("path not exist at: %s", input)
	}
	return nil
}

// ValidateIfDirExists ...
func ValidateIfDirExists(input string) error {
	if err := ValidateIfNotEmpty(input); err != nil {
		return err
	}
	if exist, err := pathutil.IsDirExists(input); err != nil {
		return fmt.Errorf("failed to check if dir exist at: %s, error: %s", input, err)
	} else if !exist {
		return fmt.Errorf("dir not exist at: %s", input)
	}
	return nil
}

// ValidateInt ...
func ValidateInt(input string) (int, error) {
	if input == "" {
		return 0, nil
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("can't convert to int, error: %v", err)
	}
	return num, nil
}

// SecureInput ...
func SecureInput(input string) string {
	if input != "" {
		return "***"
	}
	return ""
}

// ParseMultilineInput extracts each line of a multiline string separated by a new line character ("\n").
// If isVerticalBarAllowed provided "|" is also handled as a new line separator.
func ParseMultilineInput(listInput string, isVerticalBarAllowed bool) []string {
	trimmedInput := strings.TrimSpace(listInput)
	if len(trimmedInput) == 0 {
		return nil
	}

	separators := []string{"\n", `\n`}
	if isVerticalBarAllowed {
		separators = append(separators, "|")
	}

	separatedItems := []string{trimmedInput}
	for _, separator := range separators {
		separatedItems = splitElements(separatedItems, separator)
	}

	var items []string
	for _, input := range separatedItems {
		trimmedInput = strings.TrimSpace(input)
		if len(trimmedInput) > 0 {
			items = append(items, trimmedInput)
		}
	}

	return items
}

func splitElements(list []string, sep string) (s []string) {
	for _, e := range list {
		s = append(s, strings.Split(e, sep)...)
	}
	return
}
