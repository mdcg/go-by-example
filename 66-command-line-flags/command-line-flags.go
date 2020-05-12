package main
// Os "command-line flags" são formas comuns de especificar opções
// para programas de linha de comando. Por exemplo, in wc -l, o -l é
// uma "command-line flag".

// Go prove um pacote "flag" que suporta o básico de "parsing" para
// "command-line flags". Nós iremos usar esse pacote para implementa 
// o nosso exemplo.
import (
	"flag"
	"fmt"
)

func main() {
	// As declarações básicas das flags estão disponíveis para "string",
	// inteiros e booleanos. Aqui estamos declarando uma "flag" de "string"
	// com o valor padrão de "foo" e uma descrição curta. Essa "flag.String"
	// retorna um ponteiro de uma "string" (não o valor da "string" em si);
	// nós veremos como implementar esse ponteiro abaixo.
	wordPtr := flag.String("word", "foo", "a string")
	// Aqui estamos declarando outras duas flags, "fork" e "numb",
	// bastante similares à declaração acima.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	// Nós podemos também declarar uma opção que usa uma variavel já existente
	// que foi declarada em algum lugar do nosso programa. Observe que precisamos
	// passá-la como um ponteiro para a declaração da "flag".
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	// Uma vez que todas as "flags" foram declaradas, chame "flag.Parse"
	// para executar a análise da linha de comando.
	flag.Parse()
	// Aqui estamos apenas dando um "dump" nas "flags" que criamos e também para
	// qualquer argumento subsequente. Observe que precisamos desreferenciar os
	// ponteiros utilizando, por exemplo, *wordPtr para obter os valores reais
	// das opções.
	fmt.Println("word", *wordPtr)
	fmt.Println("numb", *numbPtr)
	fmt.Println("fork", *boolPtr)
	fmt.Println("svar", svar)
	fmt.Println("tail", flag.Args())
	// Para testar, a melhor forma é compilar primeiro e então executar a partir do binário.
	// $ go build command-line-flags.go

	// Um exemplo com todas as flags:
	// $ ./command-line-flags -word=opt -numb=7 -fork -svar=flag

	// Note que se você omitir alguma "flag" eles automaticamente pegam o valor padrão.
	// $ ./command-line-flags -word=opt

	// Argumentos adicionais podem ser inseridos log após as "flags".
	// $ ./command-line-flags -word=opt a1 a2 a3

	// Repare que no pacote "flag", ele requer que todas as "flags" sejam especificas antes
	// do aparecimento dos argumentos adicionais ("flags" passadas depois são consideradas
	// argumentos adicionais.
	// $ ./command-line-flags -word=opt a1 a2 a3 -numb=7

	// Use -h ou --help para gerar automaticamente um texto de ajuda no nosso programa.
	// $ ./command-line-flags -h

	// Caso você informe uma flag que não exista, o programa irá printar um erro e exibir a mensagem
	// de ajuda.
	// $ ./command-line-flags -wat
}