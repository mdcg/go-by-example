package main

// "interfaces" são denominadas coleções de assinaturas
// de métodos.
import (
	"fmt"
	"math"
)

// Aqui estamos definindo uma "interfaces" básica para
// formatos geométricos.
type geometry interface {
	area() float64
	perim() float64
}

// Para o nosso exemplo, iremos implementar essa "interfaces"
// nos tipos "rect" e "circle".
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Para implementar uma "interface" em Go, nós precisamos
// apenas implementar todos os métodos relacionados a "interface".
// Aqui nós estamos implementando os métodos de "geometry" em "rect".
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Agora a implementação parar "circles".
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Se uma variável tiver um tipo de interface, podemos chamar
// métodos que estão na interface nomeada. Aqui está uma função
// de medida (measure) genérica, que aproveita desse artífico
// parar trabalhar em qualquer geometria.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// os tipos "circle" e "rect" implementam a interface "geometry",
	// para que possamos usar instâncias dessas estruturas como argumentos
	// de "measure".
	measure(r)
	measure(c)
}
