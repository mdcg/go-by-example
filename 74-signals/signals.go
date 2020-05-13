package main
// Às vezes nós desejamos que o nosso programa em Go consiga lidar de forma inteligente
// com os Unix "signals". Por exemplo, nós podemos querer que um servidor desligue "graciosamente"
// quando ele receber um "SIGTERM", ou uma ferramente de linha de comando pare o seu processo de entrada
// quando receber um "SIGINT". Aqui mostraremos como lidar com "signals" com "channels" em Go.
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// A notificação de "signal" em Go funciona por enviar um valor "os.Signal" para um "channel".
	// Nós iremos criar um "channel" para essas notificações (nós também iremos fazer um para
	// nos notificar quando um programa puder "sair").
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	// "signal.Notify" registra o "channel" recebido para receber notificações de "signals"
	// especificos.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// Essa "goroutine" executa um recebimento bloqueante para sinais. Quando ela
	// capturar um, a mesma irá exibir e então notificar o programa que ele pode finalizar.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	// O programa irá esperar até que o "signal" esperado (como indicado pela "goroutine" acima
	// enviando um valor quando estiver pronto) e então sair.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
	// Quando nós executarmos esse programa, ele irá ficar esperando bloqueado por um "signal".
	// Pressionando "CTRL-C" (que no terminal será exibido como ^C) nós podemos enviar um "signal"
	// "SIGINT", fazendo com que o programa imprima "interrupt" e então termine.
}