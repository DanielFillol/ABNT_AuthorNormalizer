package Functions

import (
	"strings"
)

func SplitName(name string) ([]string, int) {
	text := strings.Split(name, " ")
	namesQtd := len(text)

	return text, namesQtd
}

func JuniorName(name string) bool {
	returnBool := false
	var juniorNameArray []string

	juniorNameArray = append(juniorNameArray, "filho", "filha", "neto", "neta", "junior", "júnior", "jr", "jr.", "segundo", "segunda", "terceiro", "terceira")

	for i := 0; i < len(juniorNameArray); i++ {
		if strings.ToLower(name) == juniorNameArray[i] {
			returnBool = true
		}
	}

	return returnBool
}

func Preposition(name string) bool {
	returnBool := false
	var preposicionArray []string

	preposicionArray = append(preposicionArray, "do", "da", "de", "dos", "das", "e")

	for i := 0; i < len(preposicionArray); i++ {
		if strings.ToLower(name) == preposicionArray[i] {
			returnBool = true
		}
	}

	return returnBool
}
