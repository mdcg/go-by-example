package main
// Argumentos de "Command-Line" são comuns para parametrizar a execução
// de alguns programas. Por exemplo, "go run hello.go" usa "run" e "hello.go"
// como argumentos.
import (
	"fmt"
	"os"
)

func main() {
	// os.Args provê acesso ao argumentos "brutos" da linha de comando.
	// Observe que o primeiro valor na "slice" é o caminho do programa,
	// e "args.Args[1:]" estão situados os argumentos do programa.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	// Você pode pegar elementos individuais através da indexação.
	arg := os.Args[3]
	
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	// Para testar a execução com os argumentos da linha de comando a melhor
	// forma é "buildar" o binário com o "go build".
	// $ go build command-line-arguments.go
	//	$ ./command-line-arguments a b c d
}