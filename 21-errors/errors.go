package main

// Em Go, é idiomático comunicar erros por meio de um valor de retorno
// explícito e separado. Isso contrasta com as exceções utilizadas em outras
// linguagens como Java e Ruby e o "overloaded single result", que é o valor
// de erro muitas vezes utilizando em C. Go facilita a visualização de quais funções
// retornam erros e também, a manipulação delas, utilizando a mesma estrutura de
// construção de código de tarefas que não possuem erros.
import (
	"errors"
	"fmt"
)

// Por convensão, erros são o último valor de retorno e possuem o tipo
// "error", que é uma "interface" padrão da linguagem.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constrói um valor de erro básico com a mensagem
		// passada como parâmetro.
		return -1, errors.New("can't work with 42")
	}
	// O valor "nil" na posição do retorno do erro, indica que não houve
	// nenhum erro.
	return arg + 3, nil
}

// É possível utilizar tipos customizados como erros. Para isso, você
// precisa implementar o método Error() neles. Aqui está uma variante
// que utiliza um tipo customizado para explicitamente representar
// um argumento de erro.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// Neste caso, nós utilizamos &argError para criar a nova struct,
		// fornecendo valores para os dois campos, "arg" e "prob".
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// Os dois loops abaixo testam cada uma das nossas funções de
	// retorno de erro. Observe que estamos fazendo uso do verificador
	// inline lá linha do if para checar o erro. Isso é um idioma comum
	// em Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Se você deseja usar os dados de forma programática em um
	// erro personalizado, será necessário capturá-lo com uma instância
	// do tipo personalizado do erro, por meia da asserção de tipo.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
