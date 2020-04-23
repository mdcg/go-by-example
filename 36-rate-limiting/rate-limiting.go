package main

// "Rate limiting" é um importante mecanismo para controlar utilização de recursos
// e manter a qualidade do serviço. O Go oferece elegantemente o "rate limiting" com
// "goroutines", "channels" e "tickers".
import (
	"fmt"
	"time"
)

func main() {
	// Primeiramente nós iremos ver o básico do "rate limiting". Suponha que nós queremos
	// limitar o tratamento de solicitações recebidas. Atenderemos essas solicitações em
	// um canal com o mesmo nome.
	requests := make(chan int, 5)
	for i:= 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Esse "channel" "limiter" irá receber valores a cada 200 milisegundos. Esse será o 
	// regulador no nosso esquema de "rate limiting".
	limiter := time.Tick(200 * time.Millisecond)

	// Ao bloquear um recebimento do "channel" "limiter" a cada requisição, nós limitamos
	// 1 requisição a cada 200 milissegundos.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Podemos permitir pequenas "explosões" de solicitações em nosso esquema de "rate limiting"
	// preservando o mesmo no geral. Podemos fazer isso armazenando em um "buffer" o nosso "channel"
	// "limiter". Esse "channel" "burstyLimiter" permitirá "explosões" de até três eventos.
	burstyLimiter := make(chan time.Time, 3)

	// Vamos encher o "channel" para permitir as "explosões".
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// A cada 200 milisegundos, tentamos adicionar um novo valor ao burstyLimiter, até o limite de 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	
	// Agora iremos simular mais 5 requisições recebidas. As primeiras 3 delas serão beneficiadas
	// pela capacidade de "explosão" do nosso "burtyLimiter".
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
	// Ao executar nosso programa, vemos o primeiro lote de requisições processadas de uma
	// vez a cada 200 milisegundos conforme o desejado.

	// No segundo lote de requisições nós servimos imediatamente 3, por contado da nossa 
	// "explosão" de "rate limiting", então servimos o restando das 2 com um delay de 
	// 200ms cada.
}