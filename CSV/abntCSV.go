package CSV

import (
	"fmt"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Abnt"
)

//TransformABNTCSV transform sequence of names into abnt format given by a csv
//and returns it on a given folder with the result csv
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

//createCSVs executes TransformABNT function from a []string
func createCSVs(raw []string, nameResultFolder string) error {
	var authorsABNT []Abnt.ABNTData

	for i := 0; i < len(raw); i++ {
		dataReturn, err := Abnt.TransformABNT(raw[i])
		if err != nil {
			authorsABNT = append(authorsABNT, Abnt.ABNTData{
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