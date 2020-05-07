package main

// Go oferece suporte "builtin" para "regular expressions". Aqui termos alguns
// exemplos de tarefas comuns relacionadas a esse tema em Go.
import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// Aqui estamos testando se um padrão corresponde a "string".
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	// Acima nós utilizamos um padrão de "string" diretamente, mas para
	// utilizar outra tarefa de "regexp" você precisará Compilar uma
	// otimizada "Regexp" "struct".
	r, _ := regexp.Compile("p([a-z]+)ch")
	// Vários métodos estão disponíveis nessa "struct". Aqui temos um "match"
	// como vimos anteriormente.
	fmt.Println(r.MatchString("peach"))
	// Aqui encontra uma correspondencia ao "regexp".
	fmt.Println(r.FindString("peach punch"))
	// Aqui também podemos encontrar a primeira correspondência, porém retorna
	// os indices inicial e final da correspondência, em vez da "string"
	// correspondente.
	fmt.Printn(r.FindStringIndex("peach punch"))
	// O "Submatch" inclui informações sobre as correspondências de padrão por completo
	// e as sub-correspondências na mesma. Por exemplo, aqui irá retornar informações para
	// "p([a-z]+)ch" e "([a-z]+)".
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// Similarmente, aqui irá retornar informações sobre os indices das correspondências
	// e subcorrespondências.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// A variante "All" dessas funções aplicam todas as correspondências de entrada,
	// não apenas a primeira. Abaixo segue um exemplo:
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// Essa variante de "All" estão disponíveis para outras funções como vimos acima.
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// Informando um número inteiro não-negativo como segundo argumento dessas funções
	// irá limitar o número de correspondências.
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// Nossos exemplos anteriores tiveram argumentos de "strings" e usaram nomes
	// como "MatchString". Nós também podemos usar argumentos []byte e tirarmos
	// "String" do nome da função.
	fmt.Println(r.Match([]byte("peach")))
	// Quando criamos variáveis globais com regular expressions você pode querer 
	// utilizar o "MustCompile" que é uma variação de "Compile". Basicamente a diferença entre eles
	// é que o "MustCompile" dá "panic" ao invés de retornar um erro, o que o torna mais
	// seguro para ser utilizado como variável global.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	// O pacote "regexp" também pode ser usado para dar "replace" em sublistas de strings
	// com outros valores.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	// A variante "Func" permite que vocẽ transforme um texto correspondente com uma função
	// informada.
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}