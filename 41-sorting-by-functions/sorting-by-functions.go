package main
// As vezes gostariamos de ordenar coleções por algo além
// da sua forma natural. Por exemplo, suponha que nós gostariamos
// de ordenar "strings" pelo seu tamanho ao invés da sua ordem 
// alfabética. Vamos ver um exemplo desse tipo customizado de
// ordenação em Go.
import (
	"fmt"
	"sort"
)
// Para criarmos uma função de ordenação customizada em Go,
// precisamos de um tipo correspondente. Aqui estamos criando
// um tipo byLength que é um "apelido" para o tipo "string[]" "builtin".
type byLength []string
// Nós vamos agora implementar a "interface" "sort.Interface" - "Len",
// "Less" e "Swap" - em nosso tipo para que possamos utilizar a função
// genérica "Sort" do pacote "sort". "Len" e "Swap" serão bem parecidas
// entre os dois tipos, contudo, "Less" possuirá uma lógica diferente de 
// ordenação. No nosso caso, queremos ordenar por tamanho da "string" crescente
// por isso usamos "len(s[i]) < len(s[j])" aqui.
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
// Agora que temos tudo no lugar, nós podemos implementar nosso tipo
// de ordenação customizada, convertendo o "slice" "fruits" original para
// "byLength", e então usando sort.Sort nela.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
	// Seguindo o mesmo padrão de criação de tipos customizados, implementando
	// os três métodos da "interface" em nosso tipo, e então chamando "sort.Sort"
	// na nossa coleção de tipos customizados, nós podemos ordenar os "slices" em Go
	// por funções arbitrárias.
}

