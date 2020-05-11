package main
// Ler e escrever arquivos são tarefas básicas necessárias por muitos
// programas em Go. Primeiramente vamos dar uma olhada em como ler arquivos.
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Leitura de arquivos reque a verificação da maiora das chamadas de erros.
// Este "helper" simplificará nossas verificações abaixo.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Talves a tarefa mais básica de leitura de arquivos seja
	// colocar todo o conteúdo em memória.
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))
	// Você geralmente desejará ter mais controle sobre como e quais partes
	// de um arquivo são lidas. Para essas tarefas, comece abrindo um arquivo
	// para obter um valor "os.File".
	f, err := os.Open("/tmp/dat")
	check(err)
	// Leia alguns bytes desde o início do arquivo. Permita que apenas 5 sejam
	// lidos, mas também observe quantas foram realmente lidas.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	// Você também pode procurar em um local conhecido do arquivo e começar
	// por ali.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))
	// O pacote io provê algumas funções que podem ser úteis para ler arquivos.
	// Por exemplo, leituras como as de cima podem ser implementadas com mais
	// robustez com o "ReadAtLeast".
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// Não há nenhum "builtin" para "rebobinagem", mas o "Seek"
	// realiza isso.
	_, err = f.Seek(0, 0)
	check(err)
	// O pacote "bufio" implementa um leitor em "buffer" que pode ser útil
	// tanto por sua eficiência com muitas leituras pequenas quanto pelos
	// métodos de leitura adicionais que ele fornece.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	// Feche um arquivo quando você tiver terminado de ler (geralmente você irá programar
	// logo após o "Opening" usando o "defer")
	f.Close()
	// $ echo "hello" > /tmp/dat
	// $ echo "go" >>   /tmp/dat
}
