package main
// O pacote "filepath" provê funções para você manipular e construir
// "caminhos" de foma que sejam portáveis entre sistemas operacionais;
// "dir/file" no Linux vs "dir\file" no Windows, por exemplo.
import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// O "Join" pode ser usado para construir caminhos de forma portável.
	// Ele recebe um número qualquer de argumentos e constroi um caminho
	// hierarquico a partir deles.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	// Você deve sempre utilizar o "Join" ao invés de concatenar manualmente,
	// Além de promover portabilidade, o "Join" também irá normalizar os caminhos
	// removendo qualquer separador superfluo e mudanças de diretório.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	// "Dir" e "Base" podem ser usados para "quebrar" o caminho para o diretório
	// e o arquivo. Alternativamente, "Split" irá retornar ambos na mesma chamada.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	// Cheque se um caminho é absoluto.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	// Alguns nomes de arquivos tem extensões a partir de um ponto. Nós podemos
	// "quebrar" a extensão de tal nomes com o "Ext".
	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)
	// Para encontrar o nome de um arquivo com a sua extensão removida, você pode usar
	// o "strings.TrimSuffix"
	fmt.Println(strings.TrimSuffix(filename, ext))
	// "Rel" encontra os caminhos relativos entre o "base" e o "target". Ele retorna um
	// erro caso o "target" não possa ser feito relativo ao "base".
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}