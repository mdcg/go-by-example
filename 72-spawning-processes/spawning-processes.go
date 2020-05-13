package main
// Às vezes seus programas em Go precisam criar outros processos que não
// sejam do Go. Por exemplo, no site referência, o "syntax hightlighting"
// é implementado utilizando criando um processo "pygmentize" de um programa
// em Go. Vamos ver alguns exemplos de como criar processos a partir do Go.
import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// Nós começaremos com um comando simples que não recebe argumentos ou
	// entradas, apenas imprime algo no "stdout". O "exec.Command" cria um
	// objeto que representa um processo externo.
	dateCmd := exec.Command("date")
	// ".Output" é outro auxiliar que lida com o caso comum de executar um 
	// comando, aguardar o término e coletar sua saída. Se não houverem erros,
	// o "dateOut" conterá bytes com as informações da data.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Vamos dar uma olhada em um caso um pouco mais "evoluído", em que criamos
	// um "pipe" para o processo externo em seu "stdin" e coletamos seu resultado
	// de seu stdout.
	grepCmd := exec.Command("grep", "hello")
	// Aqui nós explicitamente capturamos os "pipes" de entrada/saída, inicializamos
	// o processo, escrevemos uma entrada nele, lemos o resultado da saída, e finalmente
	// esperamos o processo terminar.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Nós omitimos a checagem de erros no exemplo acima, mas você pode user o "if err != nil"
	// para todos eles. Nós também podemos coletar os resultados de "StdoutPire", mas você pode
	// coletar o StderrPipe na exata mesma maneira.

	// Observe que, ao gerar comandos, precisamos fornecer um comando explicitamente delineado
	// e um "array" de argumentos, ao invés de passar apenas uma "string" de linha de comando.
	// Se você deseja gerar um comando completo com uma string, use a opção bash's -c.
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
	// O programa retorna saídas que são a mesma se tivessemos rodado-os diretamente da
	// linha de comando.
}