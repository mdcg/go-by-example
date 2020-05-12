package main
// Escrever um servidor HTTP é muito fácil quando utilizamos
// o pacote net/http.
import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// Um conceito fundamental no net/http são os "handlers".
	// Um "handler" é um objeto que está implementando a 
	// interface http.Handler. Uma forma comum para escrever
	// um "handler" é utilizando o adaptador "http.HandleFunc"
	// nas funções com a assinatura apropriada.

	// Funções que servem como "handlers" recebem um "http.ResponseWriter"
	// e um "http.Request" como argumentos. A resposta do "writer" é utilizada
	// para preencher a resposta HTTP. Aqui tempos nossa resposta simples que é 
	// apenas um "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// Esse "handler" faz uma coisa um pouco mais sofisticada, que é ler
	// todos os cabeçalhos da requisição HTTP e imprimi-los como o corpo
	// da resposta.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Nós registramos nossos handlers em rotas do servidor utilizando a função
	// "http.HandleFunc". Ele basicamente configura um roteador padrão no pacote
	// net/http e recebe as funções como argumento.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
	// Para rodar o servidor em background utilize o &
	// $ go run http-servers.go &
	// Para acessar a rota /hello:
	// $ curl localhost:8090/hello
}