package AbntCSV

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
	err = returnCSVABNTAuthor(raw, nameResultFolder)
	if err != nil {
		return err
	}
	fmt.Println("Files created")
	return nil
}

//returnCSVABNTAuthor executes TransformABNT function from a []string
func returnCSVABNTAuthor(raw []string, nameResultFolder string) error {
	var authorsABNT []Abnt.ABNTData

	for _, author := range raw {
		dataReturn, err := Abnt.TransformABNT(author)
		if err != nil {
			authorsABNT = append(authorsABNT, Abnt.ABNTData{
				AuthorName: author,
				ABNT:       err.Error(),
				ABNTShort:  err.Error(),
			})
		} else {
			authorsABNT = append(authorsABNT, dataReturn)
		}
	}

	err := writeCSV("filesOK", nameResultFolder, authorsABNT)
	if err != nil {
		return err
	}

	return nil
}
