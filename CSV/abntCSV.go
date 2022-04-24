package CSV

import (
	"fmt"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Abnt"
)

func TransformABNTCSV(rawFilePath string, separator rune, nameResultFolder string) error {
	raw, err := readCsvFile(rawFilePath, separator)
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
	var authorsABNT []Abnt.DataABNT

	for i := 0; i < len(raw); i++ {
		dataReturn, err := Abnt.TransformABNT(raw[i])
		if err != nil {
			authorsABNT = append(authorsABNT, Abnt.DataABNT{
				AuthorName: err.Error(),
				ABNT:       err.Error(),
				ABNTShort:  err.Error(),
			})
		}
		authorsABNT = append(authorsABNT, dataReturn)
	}

	err := writeCSV("filesOK", nameResultFolder, authorsABNT)
	if err != nil {
		return err
	}

	return nil
}
