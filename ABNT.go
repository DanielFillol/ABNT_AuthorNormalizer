package Author

import (
	"fmt"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/CSV"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Functions"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Structs"
	"strings"
)

func AbntFormat(name string) (Structs.DataABNT, error) {
	splitNames, qtdNames, err := Functions.SplitName(name)
	if err != nil {
		return Structs.DataABNT{}, err
	}

	lastName := splitNames[qtdNames-1]

	var finalName string
	var sonName string
	if Functions.JuniorName(lastName) == true {
		finalName = splitNames[qtdNames-2]
		sonName = splitNames[qtdNames-1]
	} else {
		finalName = lastName
		sonName = ""
	}

	var initialNames []string
	var initialLetters []string
	for i := 0; i < len(splitNames); i++ {
		if Functions.Preposition(splitNames[i]) == false && splitNames[i] != finalName && splitNames[i] != sonName {
			initialLetters = append(initialLetters, splitNames[i][0:1]+". ")
			initialNames = append(initialNames, splitNames[i]+" ")
		} else if splitNames[i] != finalName && splitNames[i] != sonName {
			initialNames = append(initialNames, splitNames[i]+" ")
		}
	}

	var abnt string
	if sonName != "" {
		abnt = strings.ToUpper(finalName) + " " + strings.ToUpper(sonName) + ", "
	} else {
		abnt = strings.ToUpper(finalName) + ", "
	}

	textABNTSmall := abnt
	textABNTLong := abnt

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
	}, nil
}

func AbntFormatCSV(rawFilePath string, separator rune, nameResultFolder string) error {
	raw, err := CSV.ReadCsvFile(rawFilePath, separator)
	if err != nil {
		return err
	}
	err = createCSVs(raw, nameResultFolder)
	if err != nil {
		return err
	}
	fmt.Println("Files created")
	return nil
}

func createCSVs(raw []string, nameResultFolder string) error {
	var authorsABNT []Structs.DataABNT

	for i := 0; i < len(raw); i++ {
		dataReturn, err := AbntFormat(raw[i])
		if err != nil {
			return err
		}
		authorsABNT = append(authorsABNT, dataReturn)
	}

	err := CSV.ExportCSV("filesOK", nameResultFolder, authorsABNT)
	if err != nil {
		return err
	}

	return nil
}
