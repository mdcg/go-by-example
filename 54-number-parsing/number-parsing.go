package main
// Converter números a partir de "strings" é bastante básico porém
// comum tarefa em vários programas; aqui iremos demonstrar como fazer
// isso em Go.

// O pacote "builtin" "strconv" provê tudo que precisamos
// para conversão de números.
import (
	"fmt"
	"strconv"
)

func main() {
	// Com o "ParseFloat", esse 64 diz quantos bits de precisão
	// queremos converter.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// Para "ParseInt", o 0 significa que estamos inferindo a base da "string".
	// 64 requer que esse valor seja de 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// "ParseInt" irá reconhecer formatações de números hexadecimais.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	// Um "ParseUint" também está disponível.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// Atoi é uma função conveniente para conversões de bases básicas decimais.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	// Funções de conversões retornam um erro caso entremos com alguma informação
	// errada.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}