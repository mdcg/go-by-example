package main
// A biblioteca padrão do Go veio com um excelente suporte para 
// clientes e servidores HTTP no pacote net/http. Nesse exemplo,
// nós iremos utilizar isso para demonstrar algums requisições
// HTTP simples.
import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Demonstrando uma requisição HTTP GET para o servidor. "http.Get" é
	// um atalho conveniente para criar um objeto "http.Client" e chamar
	// seu método "Get"; ele utiliza o objeto "http.DefaultClient" que
	// é uma configuração padrão muito útil.
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Aqui estamos printando o status de resposta HTTP.
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	// Vamos printar as primeiras 5 linhas do corpo da requisição. 
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}