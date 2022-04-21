package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Structs"
	"os"
	"path/filepath"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ExportCSV(nameFile string, nameFolder string, abnt []Structs.DataABNT) error {
	var authorReturn [][]string

	head := []string{"Nome Autor", "ABNT Longo", "ABNT sem ponto", "ABNT com ponto"}
	authorReturn = append(authorReturn, head)

	for i := 0; i < len(abnt); i++ {
		final := []string{
			abnt[i].AuthorName,
			abnt[i].TextABNTLong,
			abnt[i].TextABNTnoDot,
			abnt[i].TextABNTSmall,
		}
		authorReturn = append(authorReturn, final)
	}

	csvFile, _ := create(nameFolder + "/" + nameFile + ".csv")

	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	for _, newAuthor := range authorReturn {
		_ = csvWriter.Write(newAuthor)
	}

	csvWriter.Flush()

	return nil
}
