package main
// Utilize "os.Exit" para terminar imediatamente um programa
// com um status informado.
import (
	"fmt"
	"os"
)

func main() {
	// O defer >não será< executado quando utilizarmos o "os.Exit",
	// dessa forma, fmt.Println jamais será chamado
	defer fmt.Println("!")
	// Chamos o Exit com o status 3.
	os.Exit(3)
	// Note que diferente de C, Go não usa um retorno de valor inteiro
	// no main para indicar um status de "Exit". Se você quiser dar "exit"
	// com um valor diferente de zero, você deve usar o "os.Exit".

	// Quando você rodar o programa, o "exit" vai ser capturado pelo Go
	// e então printado. 
	// Apenas na hora de buildar e executar o binário você poderá ver apenas o status
	// no terminal.
}