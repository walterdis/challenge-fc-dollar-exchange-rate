package main

import (
	"io"
	"net/http"
	"os"
)

const exchangeFile string = "database/cotacao.txt"

func main() {

	request, err := http.Get("http://localhost:8081/cotacao")
	defer request.Body.Close()

	if err != nil {
		panic(err)
	}

	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	println(string(result))
	writeExchangeFile(result)
}

func writeExchangeFile(exchangeValue []byte) {
	f, err := os.Create(exchangeFile)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("Dólar: " + string(exchangeValue))
	if err != nil {
		panic(err)
	}

	f.Close()
}
