package main

// No exemplo anterior nós vimos como gerenciar contadores simples
// de estados utilizando o "atomic". Para estados mais complexos,
// nós podemos usar os "mutexes" para acessar dados de forma segura
// através de múltiplas "goroutines".
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Para o nosso exemplo o estado será um "map".
	var state = make(map[int]int)
	// Esse "mutex" irá sincronizar o nosso acesso ao estado.
	var mutex = &sync.Mutex{}

	// Nós iremos manter um rastreio de quantas operações de leitura e escrita
	// iremos fazer.
	var readOps uint64
	var writeOps uint64

	// Aqui nós iniciaremos com 100 "goroutines" para executar repetidamente leituras
	// de estado, uma por milisegundo em cada "goroutine".
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// Para cada leitura, nós iremos pegar a chave de acesso, Usar a função
				// "Lock()" do "mutex" para garantir um acesso exclusivo ao estado, ler o valor
				// na chave escolhida, usar a função "Unlock()" do "mutex", e incrementar o contador
				// "readOps".
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				// Aguardaremos um pouco entre as leituras.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Também iniciaremos 10 "goroutines" para simular escritas, utilizando o mesmo padrão
	// que fizemos para leitura.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Vamos deixar as 10 "goroutines" trabalharem no estado e no "mutex"
	// por um segundo.
	time.Sleep(time.Second)
	// Vamos ver o relatório final das operações
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// Com um último estado de "lock", vamos ver como terminou
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
	// Executando o programa nós veremos que executamos quase 90 mil operações
	// em nosso estado de "mutex" sincronizado.
}