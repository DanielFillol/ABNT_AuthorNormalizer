package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Error"
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

func ExportCSV(nameFile string, nameFolder string, abnt []Structs.DataABNT) {
	var empData [][]string

	head := []string{"Nome Autor", "ABNT Longo", "ABNT sem ponto", "ABNT com ponto"}
	empData = append(empData, head)

	for i := 0; i < len(abnt); i++ {
		final := []string{
			abnt[i].AuthorName,
			abnt[i].TextABNTLong,
			abnt[i].TextABNTnoDot,
			abnt[i].TextABNTSmall,
		}
		empData = append(empData, final)
	}

	csvFile, _ := create(nameFolder + "/" + nameFile + ".csv")
	csvWriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvWriter.Write(empRow)
	}
	csvWriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}
