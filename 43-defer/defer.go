package main
// "defer" é usado para garantir que uma chamada de função
// será feita "mais tarde" durante a execução de um programa,
// frequentemente para propóstos de "limpeza". "defer" é frequentemente
// utilizados, por exemplo, como "ensure" ou "finally" em outras linguagens.
import (
	"fmt"
	"os"
)

func main() {
	// Suponha que nós vamos criar um arquivo, lê-lo, e então fechá-lo
	// quando estivermos concluído. Aqui está um exemplo do que podemos
	// fazer com o "defer". 
	f := createFile("/tmp/defer.txt")
	// Imediatamente após instanciarmos o objeto com "createFile",
	// nós chamamos o "defer" para fechar o arquivo com "closeFile".
	// Isso irá ser executado no final da função onde está sendo chamado,
	// nesse caso, a função "main", logo após o "writeFile" ser finalizado. 
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	// É importante verificar erros quando estivermos fechando um arquivo,
	// até mesmo em uma função com "defer".
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}