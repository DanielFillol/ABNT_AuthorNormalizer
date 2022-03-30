package Author

import (
	"fmt"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/CSV"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Functions"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Structs"
	"strings"
)

func AbntFormat(name string) Structs.DataABNT {
	var abnt string
	var textABNTSmall string
	var textABNTLong string

	var finalName string
	var sonName string
	var initialNames []string
	var initialLetters []string

	splitNames, qtdNames := Functions.SplitName(name)

	lastName := splitNames[qtdNames-1]

	if Functions.JuniorName(lastName) == true {
		finalName = splitNames[qtdNames-2]
		sonName = splitNames[qtdNames-1]
	} else {
		finalName = lastName
		sonName = ""
	}

	for i := 0; i < len(splitNames); i++ {
		if Functions.Preposition(splitNames[i]) == false && splitNames[i] != finalName && splitNames[i] != sonName {
			initialLetters = append(initialLetters, splitNames[i][0:1]+". ")
			initialNames = append(initialNames, splitNames[i]+" ")
		} else if splitNames[i] != finalName && splitNames[i] != sonName {
			initialNames = append(initialNames, splitNames[i]+" ")
		}
	}

	if sonName != "" {
		abnt = strings.ToUpper(finalName) + " " + strings.ToUpper(sonName) + ", "
	} else {
		abnt = strings.ToUpper(finalName) + ", "
	}

	textABNTSmall = abnt
	textABNTLong = abnt

	for i := 0; i < len(initialLetters); i++ {
		textABNTSmall += initialLetters[i]
	}

	for i := 0; i < len(initialNames); i++ {
		textABNTLong += initialNames[i]
	}

	textABNTnoDot := strings.Replace(textABNTSmall, ".", "", -1)

	return Structs.DataABNT{
		AuthorName:    name,
		TextABNTLong:  textABNTLong,
		TextABNTnoDot: textABNTnoDot,
		TextABNTSmall: textABNTSmall,
	}
}

func DocClassifierCSV(rawFilePath string, separator rune, nameResultFolder string) {
	raw := CSV.ReadCsvFile(rawFilePath, separator)
	createCSVs(raw, nameResultFolder)
	fmt.Println("Files created")
}

func createCSVs(raw []string, nameResultFolder string) {
	var authorsABNT []Structs.DataABNT

	for i := 0; i < len(raw); i++ {
		dataReturn := AbntFormat(raw[i])
		authorsABNT = append(authorsABNT, dataReturn)
	}

	CSV.ExportCSV("filesOK", nameResultFolder, authorsABNT)

}
