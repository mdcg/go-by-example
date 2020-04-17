package main

// O exemplo anterior nós vimos como o "range" em um "for"
// provê uma iteração basica em cima de algumas estruturas de dados.
// Nós também podemos usar essa sintaxe para iterar através de valores
// recebidos por um "channel".
import "fmt"

func main() {
	// Nós iremos iterar através de 2 valores no "channel" queue
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Esse "range" itera através de cada elemento conforme é recebido de queue.
	// Como fechamos o "channel" acima, a iteração termina após o recebimento dos
	// elementos.
	for elem := range queue {
		fmt.Println(elem)
	}
	// Esse exemplo também mostra que é possível fechar um "channel" não-vazio,
	// mas ainda assim receber os valores restantes.
}
