package Abnt

import (
	"errors"
	"strings"
)

//returns an []string with all words in Test given name
func splitName(name string) ([]string, error) {
	names := strings.Split(strings.TrimSpace(name), " ")
	if len(names) < 2 {
		err := errors.New("the name contains only one word: the minimum is 2")
		return nil, err
	}

	return names, nil
}
