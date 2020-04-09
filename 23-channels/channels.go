package main

// "channels" são "pipes" que conectam "goroutines" concorrentes.
// Você pode enviar valores para canais de uma "goroutine" e recebê-los
// em outro "goroutine".
import "fmt"

func main() {
	// Crie um novo "channel" com "make(chan val-type)". "channels" são
	// tipados pelos valores que eles transmitem.
	messages := make(chan string)

	// Envie um valor para um "channel" utilizando a sintaxe "<-". Aqui
	// nós iremos enviar a mensagem "ping" para o "channel" que criamos acima,
	// a partir de uma nova "goroutine".
	go func() { messages <- "ping" }()

	// A sintaxe "<- channel" recebe o valor vindo do "channel". Aqui nós iremos
	// receber a mensagem "ping" que enviamos no trecho acima e então, iremos
	// imprimir na tela.
	msg := <-messages
	fmt.Println(msg)
	// Quando rodarmos o programa, a mensagem "ping" será transmitida com sucesso
	// de uma "goroutine" para outra via o nosso "channel".

	// Por padrão, envie e receba um bloco até que o "sender" e o "receiver"
	// estejam prontos. Essa propriedade nos permite aguardar no final do nosso programa
	// a mensagem "ping" sem precisar usar nenhuma sincronização.
}
