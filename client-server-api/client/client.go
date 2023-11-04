package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

const exchangeFile string = "../database/cotacao.txt"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/cotacao", nil)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	result, err := io.ReadAll(response.Body)
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
