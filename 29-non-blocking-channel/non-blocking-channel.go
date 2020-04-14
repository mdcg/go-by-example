package main

// Envios e recebimentos básicos em "channels" são bloqueantes.
// Entretanto, nós podemos usar o "select" com uma clausula "default"
// para implementar envios, recebimentos e até mesmo "selects multi-way"
// não bloqueantes.
import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Aqui está um recebimento não bloqueante. Se um valor
	// estiver disponível nas mensagens, o "select" selecionará
	// o caso <-messages com esse valor. Caso contrário, ele pegará
	// imediatamente o valor "default".
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Um envio não bloqueante funciona da mesma forma. Aqui o "msg"
	// não pode ser enviada para o canal de mensagens, porque o canal
	// não possui "buffer" e não há um "receiver". Portanto o caso "default"
	// está selecionado.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Podemos usar vários casos acimas do "default" para implementar
	// um "select multi-way" não bloqueante. Aqui tentamos receber sem bloqueios
	// as mensagens e os sinais.
	select {
	case msg := <-messages:
		fmt.Println("sent message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
