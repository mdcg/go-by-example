package main
// Testes unitários são uma parte importante da criação de programas em Go.
// O pacote "testing" provê algumas ferramentas importantes para nós
// escrevermos testes unitários e o comando "go test" roda os testes.

// Para fins de demonstração, esse código está no pacote "main", mas poderia
// estar em qualquer pacote. O código de teste normalmente fica no mesmo pacote
// que o código que ele testa.
import (
	"fmt"
	"testing"
)

// Nós iremos testar essa simples implementação de um "inteiro mínimo".
// Normalmente, o código que estamos testando deve ficar no arquivo fonte
// nomeado bem parecido com "intutils.go", e o teste deve ser nomeado como
// "intutils_test.go".
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// Um teste é criado escrevendo uma função que tenho o nome que comece com "Test".
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// "t.Error" irá reportar falhas no teste, mas irá continuar a execução do teste.
		// "t.Fail" irá reportar as falhas também, só que irá parar o teste imediatamente.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Escrever testes pode ser repetitivo, então é idiomático usar um estilo "table driven",
// onde as entradas de teste e saídas esperadas são listadas na tabela e um único loop irá
// percorrer e executar a lógica do teste.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d %d", tt.a, tt.b)
		// t.Run permite "subtestes", uma para cada entrada da tabela. Esses são mostrados
		// separadamentes quando executamos "go test -v".
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Para rodar os testes do projeto atual em modo "verboso":
// $ go test -v