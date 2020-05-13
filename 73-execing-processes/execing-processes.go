package main
// Nos exemplos anteriores, nós vimos como criar processos externos.
// Nós fizamos isso quando precisamos que um processo externo acessível
// a um processo Go. Às vezes nós apenas queremos trocar completamente
// o processo atual com outro (talvez um sem ser o do Go). Para fazer
// isso, nós usamos uma implementação do Go para a função clássica
// "exec".
import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Para esse exemplo nós iremos usar o "exec ls". Go requer o caminho
	// absoluto para o binário que queremos executar, então nós precisamos
	// usar o exec.LookPath para encontrá-lo (provavelmente /bin/ls).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	// "Exec" requer argumentos na forma de uma "slice" (ao contrário de uma grande
	// string). Nós vamos dar ao "ls" alguns argumentos comuns. Observe que o primeiro
	// argumento deve ser o nome do programa.
	args := []string{"ls", "-a", "-l", "-h"}

	// "Exec" também requer um conjunto de variáveis de ambiente. Aqui estamos apenas
	// passando algumas do nosso atual ambiente.
	env := os.Environ()

	// Aqui está a chamada "syscall.Exec". Se você conseguir chamá-la com sucesso, a execução
	// do nosso processo vai terminar aqui e ser substituída pelo processo "/bin/ls -a -l -h".
	// Se algum erro acontecer, nós iremos retornar o valor.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	// Após executar o programa, observe que Go não oferece a função clássica do Unix "fork".
	// Normalmente isso não é um problema, já que no início de "goroutines", a criação de processos,
	// e criação de processos a partir do "exec", abrangem a maioria dos casos de "fork".
}