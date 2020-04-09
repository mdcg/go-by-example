package main

// Nós podemos utilizar "channels" para sincronizar
// execuções entre "goroutines". Aqui faremos um exemplo
// de como um recebimento bloqueante para aguardar o término
// de uma "goroutine". Ao aguardar a conclusão de várias "goroutines"
// você pode preferir usar um WaitGroup.
import (
	"fmt"
	"time"
)

// Essa função que executaremos em uma "goroutine". O "channel"
// "done" será usado para notificar outra "goroutine" que o trabalho
// desta está concluído.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// Envia um valor para notificar que está pronto.
	done <- true
}

func main() {
	// Inicializa uma "goroutine worker", dando um "channel"
	// para notificar.
	done := make(chan bool, 1)
	go worker(done)

	// Bloqueia até receber a notificação do "goroutine worker"
	<-done

	// Se o "<-" for removido, do programa, o mesmo irá ser finalizado
	// antes mesmo do "goroutine worker" começar.
}
