package main

// Fechar um "channel" indica basicamente que nenhum valor será
// enviado a ele. Isso pode ser útil para comunicar conclusões
// aos receptores do "channel".
import "fmt"

func main() {
	// Nesse exemplo nós iremos utilizar um "channel" de "jobs"
	// para comunicar trabalhos a serem feitos da "goroutine"
	// principal para outra "goroutine" trabalhadora. Quando
	// não tivermos mais tarefas para o worker, fecharemos
	// o "channel" "jobs".
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Abaixo temos a "goroutine" worker. Ela repetidamente recebe
	// de "jobs" com "j, more := <-jobs". Nesse caso especial de recebimento
	// com dois valores, o valor "more"  será falso se "jobs" estiver fechado
	// e todos os valores no "channel" já estiverem sido recebidos. Nós usamos
	// isso para notificar no "done" quando tivermos trabalhado em todos as nossas
	// tarefas.
	go func() {
		j, more := <-jobs
		if more {
			fmt.Println("Received job", j)
		} else {
			fmt.Println("received all jobs")
			done <- true
			return
		}
	}()

	// Aqui enviamos 3 tarefas e então a fechamos.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	// Aqui nós iremos aguardar utilizando a técnica
	// de sincronização que vimos anteriormente.
	<-done
}
