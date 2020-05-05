package main

// Geralmente precisamos que nossos programas executem operações em coleções
// de dados, como selecionar todos os itens que atendem a um determinado predicado
// ou mapear todos os itens para uma nova coleção com uma função personalizada.

// Em algumas linguagens, é idiomático usar estrutas de dadaos e algoritmos
// genéricos. Go "NÃO SUPORTA" genéricos; no Go, é comum fornecer funções de
// "coleta" se e quando forem especificamente necessárias para seu programa
// e tipos de dados.

// Aqui faremos alguns exemplos de "collection functions" para "slices" e "strings".
// Você pode usar esses exemplos para construir os suas próprias funções. Observe
// que, em alguns casos, pode ser mais claro incluir o código de manipulação de coleção
// diretamente, em vez de criar e chamar uma função auxiliar.
import (
	"fmt"
	"strings"
)

// A função Index retorna o primeiro indice de uma string informada "t", ou "-1"
// caso não encontre nenhuma ocorrência.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
// Include retorna "true" se a string "t" está em um "slice".
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}
// Any retorna "true" se uma das "strings" no "slice" satisfazem o predicado "f".
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
// All retorna "true" se todos as "strings" no "slice" satisfazem o predicado "f".
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}
// Filter retorna um novo "slice" contendo apenas as "strings" no "slice" que
// satisfazem o predicado "f".
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
// Map retorna uma novo "slice" contendo o resultado da aplicação da função "f"
// para cada "string" na "slice" original.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	// Vamos testar nossas "collection functions".
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
	// Acima nós criamos várias funções anônimas, mas você pode utilizar também
	// funções nomeadas com o tipo correto.
}