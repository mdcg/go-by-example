package main

// Go oferece suporte "builtin" para XML e fomatações "XML-like"
// com o pacote encoding.xml.
import (
	"encoding/xml"
	"fmt"
)

// "Plant" será mapeada para XML. De maneiras semelhantes aos exemplos JSON,
// as tags de campos contém diretivas para o codificador e decodificador.
// Aqui usamos alguns recursos especiais do pacote XML: o nome do campo XMLName
// determina o nome do elemento XML que essa "struct" representa; id, attr significa
// que o campo Id é um atributo XML em vez de um elemento aninhado.
type Plant struct {
	XMLName xml.Name 	`xml:"plant"`
	Id		int			`xml:"id,attr"`
	Name	string		`xml:"name"`
	Origin	[]string	`xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	// Emitir XML representando nossa estrutura; usando "MarshalIndent" para
	// produzir uma saída mais legível por humanos.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	// Para adicionar um cabeçalho genérico XML para o "output", anexe-o explicitamente.
	fmt.Println(xml.Header + string(out))
	// Use Unmarshal para converter uma "stream" de bytes com XML para uma estrutura de dados.
	// Se o XML estiver "malformed" or não puder ser mapeado em "Plant", um erro bem 
	// descritivo será retornado.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	// A tag do campo parent>child>plant diz ao codificador para aninhar
	// todas as "Plants" abaixo de "<parent><child>"
	type Nesting struct {
		XMLName xml.Name	`xml:"nesting"`
		Plants	[]*Plant	`xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}