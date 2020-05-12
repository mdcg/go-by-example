package main
// Algumas ferramentas de command-line, como "go tool" ou "git" possuem
// vários subcomandos, cada um deles com suas próprias flags. Por exemplo
// "go build" e "go get" são dois subcomandos do "go tool". O pacote "flag"
// nos permite facilmente definir alguns subcomandos com suas próprias "flags".
import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Nós declaramos um subcomando utilizando a função "NewFlagSet", e 
	// prosseguimos para a definição da nova "flag" para esse subcomando.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")
	// Para diferentes subcomandos, nós podemos definir o diferentes
	// flags suportadas.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")
	// O subcomando é esperado como primeiro argumento para o programa.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	// Vamos verificar qual subcomando foi invocado.
	switch os.Args[1] {
	case "foo":
		// Para cada subcomando, nós precisamos "parsear" sua própria "flag"
		// para ter acesso aos próximos elementos adicionais.
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println(" name:", *fooName)
		fmt.Println(" tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" level:", *barLevel)
		fmt.Println(" tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	// $ go build command-line-subcommands.go 
	// Primeiramente vamos invocar um subcomando "foo":
	// $ ./command-line-subcommands foo -enable -name=joe a1 a2
	// Agora vamos tentar o "bar":
	// $ ./command-line-subcommands bar -level 8 a1
	// "bar" não aceita as flags de "foo":
	// $ ./command-line-subcommands bar -enable a1
}