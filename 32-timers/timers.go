package main

// Nós frequentemente queremos executar código Go em algum tempo
// do futuro, ou repetidamente em algum intervalo. Go possui um builtin
// "Timer" e "Ticker" fazendo que ambas as tarefas citadas anteriormente
// se tornem bem simples. Primeiramente, vamos dar uma olhada nos "timers".
import (
	"fmt"
	"time"
)

func main() {
	// Os "timers" representam um evento singular no futuro. Você diz ao
	// "timer" quanto tempo você deseja esperar, e então, ele provê um
	// "channel" que irá notificar quando esse tempo acabar. O nosso "timer"
	// definido abaixo irá esperar por 2 minutos.
	timer1 := time.NewTimer(2 * time.Second)

	// O trecho "<-timer1.C" bloqueia o "channel" do "timer" até que ele
	// envie um valor indicando que o tempo especificado acabou.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Se você apenas quisesse esperar, você poderia ter usado o time.Sleep. Um
	// motivo pelo qual o "timer" pode ser útil, é que você pode cancelá-lo antes
	// dele terminar. Aqui está um exemplo disso.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Dando tempo suficiente para o "timer2" disparar, você de fato verá que
	// ele está parado.
	time.Sleep(2 * time.Second)
	// O primeiro "timer" dispara uns ~2s após o início do programa, mas o segundo
	// deve ser interrompido antes que seja possível disparar.
}
