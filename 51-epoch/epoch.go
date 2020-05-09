package main

// Um requisito comum em programas é capturar o número de segundos, milisegundos
// ou nanosegundos desde o "Unix epoch". Aqui está como fazemos isso em Go.
import (
	"fmt"
	"time"
)

func main() {
	// Usando "time.Now" com Unix ou UnixNano para conseguir o tempo decorrido
	// desde a "Unix epoch" em segundos ou nanosegundos.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	// Note que não há "UnixMillis", então para conseguirmos os milisegundos
	// desde a época você precisará fazer a conversão manualmente.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	// Você também poderá converter segundos ou nanosegundos inteiros desde a 
	// época em seu tempo correspondente.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}