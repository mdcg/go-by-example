package main

// Go oferece um suporte extensivo para tempos e durações; aqui
// iremos ver alguns exemplos.
import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	// Vamos começar pelo tempo atual.
	now := time.Now()
	p(now)

	// Você pode construir uma "struct" "time" informando seu ano, mês, dia, etc.
	// Tempos geralmente são sempre associados com "Locations", exemplo "timezone".
	then := time.Date(20, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Você pode extrair os mais variados valores dos componentes do tempo.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// O modelo de semana de segunda a domingo também está disponível.
	p(then.Weekday())

	// Esses métodos comparam dois tempos, testando se o primeiro ocorre antes, depois ou no exato
	// momento do segundo, respectivamente.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// O método "Sub" retorna um "Duration" representando o intervalo entre dois tempos.
	diff := now.Sub(then)
	p(diff)
	// Nós podemos computar o comprimento de duração de várias unidades.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Você pode usar "Add" para avançar no tempo a partir de uma duração específica, ou
	// com um - voltar por a partir de uma duração também.
	p(then.Add(diff))
	p(then.Add(-diff))
}