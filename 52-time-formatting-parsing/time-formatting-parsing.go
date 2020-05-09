package main
// Go suporta formatação e conversão de tempo através de layouts
// "baseados em padrões".
import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	// Aqui está um exemplo básico de formatação de tempo para 
	// RFC3339, utilizando uma sua correspondente constante de layout.
	p(t.Format(time.RFC3339))
	// A conversão de tempo utiliza o mesmo valor de layout que o Format.
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	// O "Format" e "Parse" usam layouts baseados em exemplo. Geralmente,
	// você usa uma constante do tempo para esses layouts, mas também pode
	// fornecer layouts personalizados. Os layouts devem usar o "time" da
	// referência 2 de janeiro às 15:04 MST 2006 para mostrar o padrão com
	// o qual formatar/analisar um determinado "time"/"string". O "time" do
	// exemplo deve ser exatamente o mostrado: o ano de 2006, 15 a hora, segunda
	// para o dia da semana, etc.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, " 8 41 PM")
	p(t2)
	// Para representações puramente numéricas você pode utilizar também padrões
	// de formatação de "strings" com os componentes dos valores extraídos do tempo.
	fmt.Println(
		"%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	
	ansic := "Mon Jan _2 15:04:05 2006"
	// "Parse" irá retornar um erro em inputs mal formatados explicando
	// qual foi o motivo do problema.
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}