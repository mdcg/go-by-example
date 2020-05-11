package main
// Ao longo da execução de um programa, nós frequetemente criamos
// dados que não serão mais necessários depois que o programa acabar.
// Arquivos e diretórios temporários são bem úteis para esse propósito
// uma vez que eles não poluem o sistema ao longo do tempo.
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// A forma mais fácil de se criar um arquivo temporário é chamando o
	// "ioutil.TempFile". Ele cria um arquivo e o abre para leitura e escrita.
	// Nós passamos o "" como argumento, dessa forma "ioutil.TempFile" irá
	// criar um arquivo na localização padrão do nosso SO.
	f, err := ioutil.TempFile("", "sample")
	check(err)
	// Aqui vamos exibir o nome do arquivo temporário. Para sistemas 
	// "Unix-like" o diretório irá parecer com "/tmp". O nome do arquivo
	// começa com o prefixo dado como segundo argumento para "ioutil.TempFile"
	// e o resto é escolhido automaticamente para garantir que chamadas concorrentes
	// sempre irão criar arquivos com nomes diferentes.
	fmt.Println("Temp file name:", f.Name())
	// Vamos limpar tudo depois de acabarmos. O SO já é programado para limpar 
	// arquivos temporários depois de um tempo, mas ainda assim é uma boa prática
	// fazer isso explicitamente.
	defer os.Remove(f.Name())
	// Aqui nós podemos escrever algumas coisas.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	// Se você pretende criar muitos arquivos temporários, você precisa criar um
	// diretório temporário. Os argumentos de "ioutil.TempDir" são os mesmos que o 
	// "ioutil.TempFile", mas aqui ele irá retornar um nome ao invés do arquivo aberto.
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Agora nós podemos sintetizar nossos arquivos temporários com o prefixo do nosso
	// diretório temporário.
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}