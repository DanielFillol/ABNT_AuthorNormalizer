# ABNT_AuthorNormalizer
Projeto que visa transformar o nome de um autor em formato [ABNT](https://www.eco.unicamp.br/biblioteca/images/arquivos/pdf/NBR_6023__2002_-_Referencias.pdf)

## Instal
``` go get -u github.com/Darklabel91/ABNT_AuthorNormalizer ```

## Data Struct
Os dados de retorno podem ser ```bool```, ```string```, ```int``` ou ```ABNTData``` , essa última é composta por:

``` 
type ABNTData struct {
	AuthorName string `json:"AuthorName,omitempty"`
	ABNT       string `json:"abnt,omitempty"`
	ABNTShort  string `json:"abnt_short,omitempty"`
}
```

- AuthorName: Nome do autor a ser normalizado
- TextABNTLong: Nome do Autor em formato ABNT tradicional
- ABNTShort: Nome do Autor em formato ABNT tradicional com ponto 

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
Retorno
``` 
Otavio Luiz Rodrigues Júnior
RODRIGUES JÚNIOR, Otavio Luiz
RODRIGUES JÚNIOR, O. L.

Files created

 ```

## Functions

Main Function:
- TransformABNT(authorName string) -> retorna uma *ABNTData* necessitando apenas de um nome a ser processado
- TransformABNTCSV(rawFilePath string, separator rune, nameResultFolder string) -> retorna um CSV para uma pasta do projeto com o nome apontado em *resultFolder*. Para utilizar a função basta apontar o caminho do CSV (que deve conter uma única coluna com nome do autor) e o separador (';' ',' etc..)


## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
