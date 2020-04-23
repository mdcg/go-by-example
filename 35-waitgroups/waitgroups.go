package main

// Uma alternativa para aguardarmos que multiplas "goroutines"
// terminem é utilizando o wait group.
import (
	"fmt"
	"sync"
	"time"
)

// Essa é a função quie iremos rodar cada uma de nossas "goroutines". Note
// que o "WaitGroup" deve ser passado como ponteiro no argumento da função.
func worker(id int, wg *sync.WaitGroup) {
	// Ao fim da função, iremos notificar o "WaitGroup" que nós terminamos.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// O "Sleep" aqui simula uma tarefa cara.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Esse "WaitGroup" será utilizado para aguardar até que todas as nossas
	// "goroutines" terminem.
	var wg sync.WaitGroup

	// Aqui criaremos várias "goroutines" e incrementar o contador do "WaitGroup"
	// para cada uma delas.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	// Bloqueamos a execução até que o contador do "WaitGroup" volte para 0;
	// todos os workers notificaram quando eles terminaram.
	wg.Wait()
}