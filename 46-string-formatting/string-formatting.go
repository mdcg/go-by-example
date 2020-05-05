package main
// Go oferece um excelente suporte para formatação de strings 
// igual ao "printf" do C.
import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	// Go oferece vários "verbos" para usarmos na impressão que foram
	// projetados para formatar valores gerais em Go. Por exemplo, aqui
	// estamos printando uma instância da nossa "struct" "point".
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	// Se um valor é uma "struct", o "%+v" irá incluir os campos da "struct".
	fmt.Printf("%+v\n", p)
	// o "%#v" é uma variação que printa a representação da sintaxe em Go de um valor,
	// por exemplo, o snippet do código fonte que irá produzir esse valor.
	fmt.Printf("%#v\n", p)
	// Para printar o tipo de um valor, use "%T"
	fmt.Printf("%T\n", p)
	// Para formatar booleanos:
	fmt.Printf("%t\n", true)

	// Aqui temos várias opções para formatarmos inteiros.
	fmt.Printf("%d\n", 123)
	// Printar binário
	fmt.Printf("%b\n", 14)
	// Aqui printamos o caractere correspondente ao inteiro informado. 
	fmt.Printf("%c\n", 33)
	// "%x" provê o hexadecimal.
	fmt.Printf("%x\n", 456)

	// Também temos várias maneiras de printarmos "floats".
	fmt.Printf("%f\n", 78.9)
	// Aqui formatamos o "float" em uma notação parecida.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// Para printarmos strings.
	fmt.Printf("%s\n", "\"string\"")
	// Para utilizarmos as aspas duplas usamos o "%q".
	fmt.Printf("%q\n", "\"string\"")
	// Assim como nos inteiros, o "%x" renderiza uma string
	// na base hexadecimal, com dois caracteres de "outputs"
	// por byte de "input".
	fmt.Printf("%x\n", "hex this")

	// Para representarmos um ponteiro:
	fmt.Printf("%p\n", &p)

	// Quando estamos formatando números, nós frequentemente queremos
	// controlar o número do tamanho e precisão do resultado. Para 
	// especificar o tamanho de um inteiro, use um número depois do % no verbo.
	// Por padrão o resultado vai ser "justificado a direita" e ter alguns
	// paddings com espaços.
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// Você também pode especificar o tamanho dos floats printados, embora
	// geralmente você deseje restringir a precisão decimal ao mesmo tempo com a
	// sintaxe "tamanho.precisão"
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// Para justificar para a esquerda, use o a flag "-".
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Você também pode controlar o tamanho quando estiver trabalhando com "strings",
	// especialmente para garantir que eles estaram alinhados bem parecidos com "tabelas".
	// Assim justificamos eles a direita.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// Para justificarmos a esquerda, usamos a flag "-".
	fmt.Printf("|%-6s|%-6s\n", "foo", "b")

	// Até então nós vimos o "Printf", que imprime a string formatada em "os.Stdout". O 
	// Sprintf formata uma sequência sem imprimi-la em qualquer lugar.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	// Você pode também formatar+printar para "io.Writers" diferentes de os.Stdout usando
	// o Fprintf.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}