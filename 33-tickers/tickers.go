package main

// "Timers" são para quando nós queremos fazer alguma coisa no futuro.
// "Tickers" são para quando nós queremos fazer alguma coisa repetidamente
// em intervalos regulares.
import (
	"fmt"
	"time"
)

func main() {
	// "Tickers" usam um mecanismo similar aos "timers": Um "channel" que irá
	// enviar valores. Aqui nós utilizaremos o "select" em um "channel" para 
	// esperar assincronamente valores que chegaram a cada 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <- done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// "Tickers" também podem ser parados como "timers". Uma vez que ele esteja parado,
	// nós não receberemos mais valores nesse "channel". Iremos parar o nosso "ticker"
	// após 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
	// Quando rodarmos o programa, veremos que o "ticker" executa uma três vezes antes de parar.
}