package main

// Go suporta métodos definidos em tipos "struct".
import "fmt"

type rect struct {
	width, height int
}

// Esse método "area" possui um "tipo receptor" de *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Métodos podem ser definidos tanto para tipos receptores de ponteiros
// quanto de valores. Abaixo segue um exemplo de um receptor de valor.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Nós podemos chamar os dois métodos definidos para a nossa "struct".
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go lida automaticamente com a conversão entre valores e ponteiros
	// para chamadas de método. Você pode usar um tipo receptor de ponteiro
	// para evitar chamadas de método ou permitir que um método mude a estrutura
	// do recebimento.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
