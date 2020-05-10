package main
// Hashes SHA1 são frequentemente utilizados para calcular identidades
// curtas para "blobs" binária ou de texto. Por exemplo, o sistema de controle
// de versão "git" usa SHA1s extensivamente para identificar arquivos e diretórios
// com versão. Veja como calcular hashes SHA1 no Go.

// Go implementa várias funções de "hash" em vários pacotes crypo/*.
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"
	// O padrão para gerar um hash é "sha1.New()", "sha1.Write(bytes)", e então
	// sha1.Sum([]bytes{}). Aqui nós estamos começando um novo hash.
	h := sha1.New()
	// "Write" espera bytes. Se você tem uma "string" s, use []bytes(s) para convertê-la
	// em bytes.
	h.Write([]byte(s))
	// Isso aqui capturar a hash finalizada como um "slice" de bytes. O argumento
	// para "Sum" pode ser utilizado para concatenar a uma "slice" existente: geralmente
	// isso não é necessário.
	bs := h.Sum(nil)
	// Valores SHA1 são frequentemente printados como hexadecimais, por exemplo, em commits do "git".
	// Use o formato "%x" para converter o resultado hash para uma "string" hexadecimal.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	// Executando o programa irá computar o hash e printá-lo de uma forma hexidecimal legíveis para
	// seres humanos.

	// Você pode computar outros hashes usando padrões similares para o exibido na execução do programa.
	// Por exemplo, para computar hashes MD5 importamos crypto/md5 e use md5.New().
}