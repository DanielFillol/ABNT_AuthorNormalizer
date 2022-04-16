# ABNT_AuthorNormalizer
Projeto que visa transformar o nome de um autor em formato ABNT

## Instal
``` go get -u github.com/Darklabel91/ABNT_AuthorNormalizer ```

## Data Struct
Os dados de retorno podem ser ```bool```, ```string```, ```int``` ou ```DataABNT``` , essa última é composta por:

``` 
type DataABNT struct {
	AuthorName    string `json:"AuthorName,omitempty"`
	TextABNTLong  string `json:"TextABNTLong,omitempty"`
	TextABNTnoDot string `json:"TextABNTnoDot,omitempty"`
	TextABNTSmall string `json:"TextABNTSmall,omitempty"`
}
```

- AuthorName: Nome do autor a ser normalizado
- TextABNTLong: Nome do Autor em formato ABNT tradicional
- TextABNTnoDot: Nome do Autor em formato ABNT tradicional com ponto 
- TextABNTSmall: Nome do Autor em formato ABNT somente iniciais

## Example

``` 
package main

import (
	"fmt"
	Author "github.com/Darklabel91/ABNT_AuthorNormalizer"
)

func main() {

	//SINGLE USE OF THE CLASSIFIER

	authorName := "Otavio Luiz Rodrigues Júnior"

	ABNT := Author.AbntFormat(authorName)

	fmt.Println(ABNT.TextABNTLong)
	fmt.Println(ABNT.TextABNTnoDot)
	fmt.Println(ABNT.TextABNTSmall)

	//READING A CSV WITH AUTHOR'S

	rawPath := "/Users/Desktop/Authors.csv"
	separator := ','
	resultFolder := "Result"

	Author.AbntFormatCSV(rawPath, separator, resultFolder)

}

 ```
Retorno
``` 
RODRIGUES JÚNIOR, Otavio Luiz 
RODRIGUES JÚNIOR, O L 
RODRIGUES JÚNIOR, O. L. 
Files created

 ```

## Functions

Main Function:
- AbntFormat(name string) (string, string, string)  -> retorna uma *DataABNT* necessitando apenas de um nome a ser processado
- AbntFormatCSV(rawFilePath string, separator rune, nameResultFolder string) -> retorna um CSV para uma pasta do projeto com o nome apontado em *resultFolder*. Para utilizar a função basta apontar o caminho do CSV (que deve conter uma única coluna {authorName}) e o separador (';' ',' etc..)

Suport Functions:
- SplitName(name string)    ->  retorna um *array de string* e *int* sendo os *nomes* e a *quantidade de nomes existentes*
- JuniorName(name string)   ->  retorna bool para um nome identificado como *júnior, filho, filha, segundo, terceiro, etc.*
- Preposition(name string)  ->  retorna bool para um nome identificado como *de, do, das, dos, e, etc.*

## Considerações
A) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.

B) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
