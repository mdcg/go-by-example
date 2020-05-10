package main
// URLs provê uma forma uniforme de localizar recursos. Aqui mostraremos
// como converter URLs em Go.
import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// Nós iremos converter esse exemplo de URL, que inclui um "scheme", informações de autenticação,
	// "host", "port", "path", parâmetros de consultas, e fragmentos de consultas.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// Converta uma URL e garanta que não há nenhum erro.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// Acessar um "scheme" é direto.
	fmt.Println(u.Scheme)
	// User contém todas as informações de autenticação; chamando "Username" e "Password"
	// dessa forma para ter os valores individuais respectivos.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	// O "Host" contém ambos o "hostname" e a porta, se estiver presente. Utilize "SplitHostPort"
	// para extraí-los.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	// Aqui nós estamos extraindo o "path" e o "fragmento" depois do #.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	// Para capturar os "query params" em uma "string" de no formato k=v, use "RawQuery". 
	// Você pode também converter "query params" em "maps". Os "maps" de "query params"
	// são "strings" de "slices" de "strings", então no index em [0] você irá chamar se quiser
	// o primeiro valor.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}