package main
// Go prove suporte "builtin" para codificação/decodificação base64.

// Essa sintaxe importa o pacote "encoding/base64" com o nome "b64"
// ao invés do padrão "base64". Isso irá nos poupar alguns espaçamentos
// abaixo.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// Aqui temos uma "string" que queremos codificar/decodificar.
	data := "abc123!?$*&()'-=@~"
	// Go suporta ambos os padrões e "URL-compatible" base64. Aqui nós
	// mostramos como codificar utilizando o padrão "encoder". O "encoder"
	// requer um []byte então nós precisamos converter nossa "string" para
	// esse tipo.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	// Ao decodificarmos talvez tenhamos um erro, que poderá checar se
	// a entrada está bem formatada.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	// Aqui codificamos/decodificamos usando um "URL-compatible" formato base64.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
	// "String" codifica para valores ligeiramente diferentes com os codificadores padrão
	// e URLs base64 (a direita + vs -), mas ambos decodificam para a cadeia original,
	// conforme desejado.
}