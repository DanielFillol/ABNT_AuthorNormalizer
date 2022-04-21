package Functions

import (
	"errors"
	"strings"
)

func SplitName(name string) ([]string, int, error) {
	text := strings.Split(name, " ")
	namesQtd := len(text)
	if len(text) < 2 {
		err := errors.New("there are no \" \" in the name")
		return nil, 0, err
	}

	return text, namesQtd, nil
}

func JuniorName(name string) bool {
	juniorNameArray := []string{"filho", "filha", "neto", "neta", "junior", "júnior", "jr", "jr.", "segundo", "segunda", "terceiro", "terceira"}

	for _, junior := range juniorNameArray {
		if strings.ToLower(name) == junior {
			return true
		}
	}

	return false
}

func Preposition(name string) bool {
	prepositionArray := []string{"do", "da", "de", "dos", "das", "e"}

	for _, prepArray := range prepositionArray {
		if strings.ToLower(name) == prepArray {
			return true
		}
	}

	return false
}
