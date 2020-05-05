package main
// Da biblioteca padrão, o pacote "strings" provê várias funções super úteis
// para tarefas relacionadas a "string". Vamos agora fazer alguns 
// exemplos para te dar um norte na utilização desse pacote.
import (
	"fmt"
	s "strings"
)
// Nós apelidamos o "fmt.Println" para encurtar o nome já que vamos utilizá-lo
// bastante nos exemplos abaixo.
var p = fmt.Println

func main() {
	// Aqui temos os os exemplos das funções pertencentes ao pacote "strings".
	// Como essas funções são do pacote, não métodos no próprio objeto, precisamos
	// passar a "string" como o primeiro argumento da função. Você pode ler mais sobre
	// a utilização do pacote na documentação oficial.
	p("Contains:	", s.Contains("test", "es"))	
	p("Count:		", s.Count("test", "t"))
	p("HasPrefix:	", s.HasPrefix("test", "te"))
	p("HasSuffix:	", s.HasSuffix("test", "st"))
	p("Index:		", s.Index("test", "e"))
	p("Join:		", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:		", s.Repeat("a", 5))
	p("Replace:		", s.Replace("foo", "o", "0", -1))
	p("Replace:		", s.Replace("foo", "o", "0", 1))
	p("Split:		", s.Split("a-b-c-d-e", "-"))
	p("ToLower:		", s.ToLower("TEST"))
	p("ToUpper:		", s.ToUpper("test"))
	p()
	// Esse trecho não faz parte do pacote "strings", mas vale a pena mencionarmos que abaixo
	// estamos utilizando alguns mecanismos para obtermos os valores do tamanho "string" em bytes
	// e também, obtermos um byte através de seu indice.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
	// Note que "len" and "indexing" funcionam em "byte-level". Go usa o "UTF-8 encoded" para "strings"
	// então isso é frequentemente útil. Se você potencialmente tiver que trabalhar com caractéres
	// "multi-byte", você vai precisar usar operações de "encoding-aware".
}