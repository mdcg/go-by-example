package main

// No exemplo anterior nós utilizamos um "locking" explícito com 
// "mutexes" para sincronizar o acesso para estados compartilhados
// entre várias "goroutines". Outra opção é utilizar o builtin de 
// sincronização de "goroutines" e "channels" para chegarmos ao mesmo
// resultado. Esse abordagem "channel-based" alinhas as ideias de Go
// de compartilhar memória comunicando e possuindo cada parte de dados
// de exatamente uma (1) "goroutine".
import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Nesse exemplo nosso estados pertencerão por uma única "goroutine".
// Isso irá garantir que o dado nunca será corrompido com acessos concorrentes.
// Para ler ou escrever esse estado, outras "goroutines" enviarão mensagens
// para a "goroutine" proprietária e receberão respostas concorrentes. Essas
// "structs" "readOp" e "writeOp" encapsulam essas solicitações e uma maneira
// da "goroutine" proprietária responder.
type readOp struct {
	key 	int
	resp 	chan int
}

type writeOp struct {
	key 	int
	val 	int
	resp 	chan bool
}

func main() {
	// Como no exemplo anterior, nós vamos ler quantas
	// operações de leitura e escrita serão realizadas.
	var readOps uint64
	var writeOps uint64

	// Esses canais serão usados por outras "goroutines"
	// para emitir solicitações de leitura e escrita.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Aqui está a "goroutine" que possui o estado, que é um "map" como no exemplo anterior
	// mas agora privado para a "stateful goroutine". Esta "goroutine" seleciona repetidamente
	// nos "channels" de leitura e gravação, respondendo às solicitações assim que chegam. Uma resposta
	// é executada executando primeiro a operação solicitada e depois enviando um valor no "channel"
	// de resposta "resp" para indicar sucesso (e o valor desejado no caso de leituras).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Aqui iniciaremos 100 "goroutines" para emitir leituras para o "state-owning goroutine" via
	// "channel" "reads". Cada leitura requer inicializar um readOp, enviando ele através do "channel"
	// "reads" e recebendo o resultado através do "channel" "resp".
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:	rand.Intn(5),
					resp:	make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Aqui iniciaremos 10 "goroutines", com uma abordagem similar.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:	rand.Intn(5),
					val:	rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Vamos deixar as "goroutines" trabalharem por um segundo.
	time.Sleep(time.Second)

	// Finalmente iremos capturar e relatório de contagem de operações.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// Rodando nosso programa, vemos que nosso gerenciamento de estado com "goroutines"
	// (goroutine-based state management) completa aproximadamente 80.000 operações.

	// Para esse caso em particular, a abordagem baseada em "goroutines" foi um pouco mais envolvida
	// do que a baseada em "mutex". Contudo, pode ser útil em alguns casos, por exemplo, onde você
	// tem outros "channels" envolvidos ou quando você gerencia vários desses "mutexes", eles podem estar
	// sujeitos a erros. Você deve usar a abordagem que parecer mais natural, principalmente no que 
	// diz respeito à compreensão do programa.
}