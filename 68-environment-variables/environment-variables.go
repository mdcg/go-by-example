package main
// Variáveis de ambiente são mecanismos universais para transmitir
// informações de configuração para programas Unix. Vamos dar uma olhada
// em como "setar", capturar e listar as variáveis de ambiente.
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Para "setar" um par "chave/valor", use "os.Setenv". Para capturar
	// uma chave use "os.Getenv". Isso irá retornar uma string vazia se
	// a chave não estiver presente nas variáveis de ambiente.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// Use "os.Environ" para listar todas os pares "chave/valor" no ambiente.
	// Isso irá retornar um "slice" de "strings" na forma "KEY=value". Você
	// pode utilizar "strings.SplitN" e então capturar a chave e o valor. Aqui
	// estamos exibindo todas as variáveis de ambiente.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
	// Ao rodarmos o nosso programa, vamos observer que obtivemos o valor de "FOO"
	// contudo "BAR" está vazio.
	// $ go run environment-variables.go
	// A lista das variáveis de ambiente vai ser diferente de sistema pra sistema.
	// Se setarmos a variável "BAR" no ambiente primeiro, então ao executarmos o 
	// programa teremos o seu valor correspondente.
}