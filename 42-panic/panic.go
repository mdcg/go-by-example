package main
// Um "panic" significa tipicamente que alguma inesperada
// aconteceu de errado. Mais frequentemente, o usamos para
// falhar rapidamente em erros que não devem ocorrer durante
// uma operação normal, ou que não estamos preparados para lidar
// com "graciosidade".
import "os"

func main() {
	// Nós iremos o "panic" em todo esse "site" para checar por
	// erros inesperados. Ese é o único programa no "site"
	// projetado para entrar em "pânico".
	panic("a problem")
	// Um uso comum do "panic" é para abortar se uma função retorna
	// um valor de "error" que nós não sabemos como (ou não queremos)
	// lidar. Um exemplo de "panicking" se nós capturarmos um erro
	// inesperado ao criamos um novo arquivo.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
	// Executando nosso programa vamos causar o seu "panic", que irá printar
	// um erro e todo o rastreio da "goroutine" e então, sair com um status diferente
	// de zero.

	// Note que diferente de outras linguagens que usamos exceções para manipularmos muitos erros,
	// em Go é idiomático usar valores de retorno "error-indicating" sempre que possível.
}