package main

// Nesse exemplo, iremos ver como implementar um "worker pool"
// utilizando "goroutines" e "channels".
import (
	"fmt"
	"time"
)

// Aqui estamos definindo o nosso "worker" no qual executaremos
// várias instâncias simultaneamentes. Esses "workers" receberão
// tarefas no "channel" "jobs" e enviarão os resultados correspondentes
// para "results". Dormiremos um segundo por trabalho para simular uma
// tarefa "custosa".
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "stated job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// Para utilizar nosso conjunto de "workers", precisamos
	// enviar tarefas e coletar seus resultados. Para isso,
	// iremos utilizar dois "channels".
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Esse trecho irá inicializar três "workers", inicialmente estarão
	// bloqueados porque não temos tarefas ainda.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// Aqui iremos enviar 5 tarefas e então fechar o "channel" para indicar
	// que todo trabalho terminou.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Por fim, nós iremos coletar os resultados das tarefas. Isso garante que as "goroutines"
	// tenham terminado. Uma maneira de aguardar várias "goroutines" é usar um "WaitGroup".
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	// Ao executarmos o nosso programa, veremos que 5 tarefas começaram a ser executadas
	// por vários "workers". O programa demora apenas cerca de 2 segundos ao invés de demorar
	// 5 segundos no total porque temos três workers rodando paralelamente.
}