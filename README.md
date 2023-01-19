# ABNT_AuthorNormalizer
Packge to transform authors of papers and books in [ABNT](https://www.eco.unicamp.br/biblioteca/images/arquivos/pdf/NBR_6023__2002_-_Referencias.pdf) (Brazilian Association of Technical Standards) format.

## Instal
``` go get -u github.com/Darklabel91/ABNT_AuthorNormalizer ```

## Data Struct
Return data can be: ```bool```, ```string```, ```int``` or ```ABNTData```, the last
``` 
type ABNTData struct {
	AuthorName                string `json:"AuthorName,omitempty"`
	ABNT                      string `json:"abnt,omitempty"`
	ABNTShort                 string `json:"abnt_short,omitempty"`
	FirstLetters              string `json:"abnt_firstLetters,omitempty"`
	FirstLettersButCompanySig string `json:"abnt_firstLettersButCompanySig,omitempty"`
}
```

- AuthorName: name to be normalized
- TextABNTLong: name in ABNT format
- ABNTShort: name on simple ABNT format
- FirstLetters: only the initials of a name regardless of company names
- FirstLettersButCompanySig: only the initials without company names

## Example
``` 
package main

import (
	"fmt"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Abnt"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/AbntCSV"
)

func main() {

	//Single use
	authorName := "Otavio Luiz Rodrigues Júnior"

	test, err := Abnt.TransformABNT(authorName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(test.AuthorName)
	fmt.Println(test.ABNT)
	fmt.Println(test.ABNTShort)

	//CSV use
	//path := "/Users/Desktop/Authors.csv"
	path := "/Users/danielfillol/Documents/GitHub/LegalDoc_Classifier/ABNT_AuthorNormalizer/CSV/authors.csv"
	separator := ','
	nameResultFolder := "test"

	err = AbntCSV.TransformABNTCSV(path, separator, nameResultFolder)
	if err != nil {
		fmt.Println(err)
	}
}
 ```
Return
``` 
Otavio Luiz Rodrigues Júnior
RODRIGUES JÚNIOR, Otavio Luiz
RODRIGUES JÚNIOR, O. L.
O. L. R. J.

Files created

 ```

## Functions

Main Function:
- TransformABNT(authorName string) -> return *ABNTData* for a given name.
- TransformABNTCSV(rawFilePath string, separator rune, nameResultFolder string) -> return CSV to the given folder on param *resultFolder*. The .csv file given on param *rawFilePath* must contain only a single columm with names.a


