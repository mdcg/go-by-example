package main

// O mecanismo primário para gerenciar estados em Go é a comunicação
// entre "channels". Nós vimos esse exemplo utilizando os "workers pools"
// Existem algumas outras opções para gerenciar estado também. Nós iremos
// utilizar o pacote "sync/atomic" para "atomic counters" acessados por
// múltiplas "goroutines".
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Nós iremos utilizar um inteiro sempre positivo para representar nosso contador.
	var ops uint64
	// Um "WaitGroup" irá nos ajudar a esperar por que todas as "goroutines" terminem suas
	// tarefas.
	var wg sync.WaitGroup

	// Nós iremos iniciar 50 "goroutines" onde cada uma irá incrementar o nosso contador exatamente
	// 1000 vezes.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// Para incrementar "atomicamente" o contador, nós usaremos "AddUint64", informando
				// o endereço de memória do nosso contador "ops" com a sintaxe &.
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
			}()
		}
	// Aguarde até que as "goroutines" terminem.
	wg.Wait()
	// Agora é seguro acessar operações porque sabemos que nenhuma outra "goroutine" está escrevendo
	// nela. Também é possível ler "atômicos" com segurança enquanto estão sendo atualizados, usando
	// funções como "atomic.LoadUint64".
	fmt.Println("ops:", ops)
	// Esperamos obter exatamente 50.000 operações. Se tivéssemos usado as operações não atômicas
	// para incrementar o contador "ops" (++), provavelmente obteríamos um número diferente, alternando
	// entre execuções, porque goroutines interfeririam umas nas outras. Além disso, teremos falhas no
	// "data race" ao executar com o sinalizador "-race".
}