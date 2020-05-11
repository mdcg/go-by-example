package main
// Escrever arquivos em Go segue padrões similares ao que vimos mais
// cedo para leitura de arquivos.
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Para começar, aqui demonstraremos como "dumpar" um "string" (ou apenas bytes)
	// em um arquivo.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)
	// Para escritas mais granulares, abra o arquivo para escrita.
	f, err := os.Create("/tmp/dat2")
	check(err)
	// É idiomático usar o "defer" "Close" para fechar imediatamente após a abertura
	// de um arquivo.
	defer f.Close()
	// Você pode escrever utilizando "slice" de bytes.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	// Um "WriteString" também está disponível.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	// Emita um "Sync" para "liberar" gravações no armazenamento.
	f.Sync()
	// "bufio" provê escritores "bufferizados" assim como os leitores que vimos anteriormente.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	// Utilize "Flush" para garantir todas as operações em buffer tenham sido aplicadas ao
	// escritor.
	w.Flush()
}