package main
// No exemplo anterior nós vimos como configurar um simples servidor HTTP. Os servidores HTTP
// são úteis para demonstrar a utilização do "context.Context" por controlar o cancelamento. O
// "Context" possui "deadlines", "cancellation signals", e outros valores no escopo da requisição
// através dos limites da API e das "goroutines".
import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// Um "context.Context" é criado para cada requisição pelo mecanismo "net/http" e está disponível
	// com o método "Context".
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")
	// Espera alguns segundos antes de enviar ou responder um cliente. Isso pode simular algum trabalho
	// que está ocorrendo. Enquanto trabalhamos, mantenha os olhos no "channel" "context.Done()" para um
	// "signal" que deve ser de cancelamento da tarefa e retorno o mais rápido possível.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// O método "context.Err()" retorna um error que explica porque o "channel" "Done()" foi fechado.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// Como antes, vamos registrar nosso "handler" na rota "/hello" e então iniciarmos o servidor.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
	// Para rodar o server em background:
	// $ go run context.go &
	// Simule uma requisição para /hello, então pressione CTRL + C um pouco depois de iniciar
	// para emitir o "signal" de cancelamento.
}