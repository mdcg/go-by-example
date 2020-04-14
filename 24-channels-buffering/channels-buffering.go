package main

// Por padrão, os "channels" não são armazenados em
// "buffers", o que significa que eles só aceitarão
// envios (chan <-) se houver recebimento correspondente
// (<- chan) pronto para receber o valor enviado. os "channels"
// em "buffer" aceitam um número limitado de valores em um
// receptor correspondente para eles.
import "fmt"

func main() {
	// Vamos criar um "channel" de strings que irá "bufferizar"
	// no máximo dois valores.
	messages := make(chan string, 2)

	// Como esse "channel" é "bufferizado", podemos enviar
	// esses valores para um "channel" sem um recebimento
	// concorrente correspondente.
	messages <- "buffered"
	messages <- "channel"

	// Mais tarde, podemos receber esses dois valores como
	// de costume.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
