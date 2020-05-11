package main
// Go possui diversas funções para trabalharmos com diretórios
// no file system.
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
	// Vamos criar um subdiretório no diretório atual.
	err := os.Mkdir("subdir", 0755)
	check(err)
	// Quando criamos diretórios temporários, é uma boa prática
	// em "defer" sua remoção. "os.RemoveAll" irá deletar toda
	// árvore de diretório (similar ao rm -rf).
	defer os.RemoveAll("subdir")
	// Aqui temos uma função que irá nos ajudar a criarmos
	// diretórios vazios.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Nós podemos criar um hierarquia de diretórios, incluindo pais
	// com "os.MkdirAll". Isso é similar ao comando mkdir -p.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// "ReadDir" lista o conteúdo do diretório, retornando um "slice"
	// de objetos do tipo "os.FileInfo".
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// "Chdir" nos permite mudarmos de diretório, similar ao cd.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Agora vamos o conteúdo de "subdir/parent/child" quando listarmos
	// o diretório atual.
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Vamos dar um cd até onde nós começamos.
	err = os.Chdir("../../..")
	check(err)

	// Nós podemos também visitar diretórios recursivamente, incluindo todos
	// os subdiretórios. "Walk" aceita uma função de "callback" para lidar
	// com cada arquivo ou diretório visitado.
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// Essa função é chamada para cada arquivo ou diretório encontrado recursivamente
// pelo filepath.Walk.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}