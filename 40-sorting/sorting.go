package main
// Os pacotes de ordenação de Go implementa funções
// tanto para tipos "builtin" quanto para outros
// tipos definidos pelo usuário. Vamos primeiramente
// usá-los para ordenação de tipos "builtin"
import (
	"fmt"
	"sort"
)

func main() {
	// Os métodos de "sort" são específicos para "builtins";
	// aqui está um exemplo para strings. Repare que a ordenação
	// não retorna um novo tipo, ela ordena a estrutura informada
	// no ato.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	// Analogamente, também podemos ordenar inteiros
	ints := []int{7, 2, 4}
	sort.ints(ints)
	fmt.Println("Ints:", ints)
	// Podemos também checar se um "slice" estão ordenados. Essa
	// função retorna um valor booleano.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}