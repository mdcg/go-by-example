package main

// Um "line filter" é um tipo comum de programa que lê uma entrada no "stdin",
// processa, e então imprimi alguma valor derivado no "stdout". "grep" e "sed"
// são "line filters", por exemplo.

// Aqui faremos um exemplo de "line filters" em Go que irá deixar em maiúsculo
// todas as entradas de texto. Você pode usar esse padrão para escrver seus
// próprios "line filters".
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Pulando para o sem buffer os.Stdin com um scanner bufferizado, temos um método
	// conveniente de digitalização que avança o scanner para o próximo token; no qual
	// é a próxima linha no scanner padrão.
	scanner := bufio.NewScanner(os.Stdin)
	// "Text" retorna o token atual, aqui na próxima linha, da entrada.
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		// Aqui iremos imprimir a linha em caixa alta.
		fmt.Println(ucl)
	}
	// Verifique se há erros durante o Scan. O fim da linha é esperado e não é reportado
	// pelo scan como um erro.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	// $ echo 'hello'   > /tmp/lines
	// $ echo 'filter' >> /tmp/lines
	// $ cat /tmp/lines | go run line-filters.go
}