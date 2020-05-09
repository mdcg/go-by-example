package main

// Go oferece alguns "builtins" de suport para codificar de descoficiar JSON,
// incluindo e para tipos de dados "builtin" e personalizados.
import (
	"encoding/json"
	"fmt"
	"os"
)

// Nós iremos utilizar duas "structs" para demonstrar a codificação e 
// descodificação de tipos personalizados abaixo.
type response1 struct {
	Page 	int
	Fruits 	[]string
}
// Apenas campos exportados serão codificados/descodificados em JSON.
// Campos precisam começar com letras maiúsculas para serem exportáveis.
type response2 struct {
	Page 	int			`json:"page"`
	Fruits	[]string	`json:"fruits"`
}

func main() {
	// Primeiramente nós iremos ver a codificação de dados básicos
	// para "strings" JSON. Aqui está alguns exemplos para valores atômicos.
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Aqui nós temos alguns "slices" e "maps", que serão codificados
	// para arrays e objetos de JSON.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// O pacote JSON pode automaticamente codificar tipos personalizados
	// de dados. Ele irá apenas incluir os campos exportáveis para a saída
	// da codificação e por padrão ele utilizará o nomes dos campos
	// como chaves do JSON.
	res1D := &response1{
		Page: 	1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Você pode utilizar "tags" nas declarações dos campos da "struct" para
	// personalizar os nomes das chaves do JSON que foram codificados.
	// Veja a declaração de response2 acima para ver os exemplos das "tags".
	res2D := &response2{
		Page:	1,
		Fruits:	[]string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(res2B)

	// Agora iremos dar uma olhada na descodificação de dados em JSON para
	// valores em Go. Aqui está um exemplo de uma estrutura de dados genérica.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// Nós precisamos prover uma variavel aonde o pacote JSON irá colocar os dados
	// descodificados. Esse "map[string]interface{}" irá nos dar um "map" de "strings"
	// para valores arbitrários de tipos de dados.
	var dat map[string]interface{}
	// Aqui que está realmente a etapa para descoficiação e checagem de erros durante
	// o processo.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Para usar os valores no "map" decodificado, precisamos convertê-lo para o tipo apropriado.
	// Por exemplo, aqui convertemos o valor em "num" para o tipo esperado "float64".
	num := dat["num"].(float64)
	fmt.Println(num)
	// O acesso a dados aninhados requer uma série de conversões.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Nós também podemos decodificar JSON para tipos de dados personalizados. Isso tem as vantagens
	// de adicionar "segurança de tipo" adicional aos nossos programas e eliminar a necessidade
	// de asserções de tipo ao acessar os dados decodificados.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// Nos exemplos acima, sempre usamos "bytes" e "strings" como intermediários entre os dados e a representação
	// JSON na saída padrão. Também podemos transmitir codificações JSON diretamente para os.Writers como os.Stdout ou
	// mesmo para corpos de resposta HTTP.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}