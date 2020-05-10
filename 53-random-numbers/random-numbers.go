package main

// O pacote math/rand de Go provê a geração de alguns 
// números pseudorandômicos.
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Por exemplo, "rand.Intn" retorna um inteiro n, 0 <= n <= 100.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	// "rand.Float64" retorna um float64 f, 0.0 <= f <= 1.0.
	fmt.Println(rand.Float64())
	// Isso pode ser utilizado para gerar floats em outros intervalos,
	// por exemplo, 5.0 <= f <= 10.0.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// O número padrão do gerador é deterministico, então ele irá produzir
	// a mesma sequência de número cada vez por padrão. Para produzir sequencias
	// variantes, dê a ela uma "seed" que muda. Note que isso não é muito seguro
	// para números aleátorios que você deseja que sejam secretos; use "cripto/rand"
	// para esses casos.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// Chamando o resultante "rand.Rand" como essas funções no pacote "rand".
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	// Se você quer "seedar" a fonte com o mesmo número, isso irá produzir a mesma
	// sequência de números aleatórios.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
	// Veja a documentação do pacote math/rand para outros tipos de valores
	// randomicos que o Go pode prover.
}