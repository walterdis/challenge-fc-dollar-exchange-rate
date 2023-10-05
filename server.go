package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/walterdis/challenge-fc-dollar-exchange-rate/src"
	"io"
	"net/http"
	"time"
)

const dollarApiUrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func main() {
	http.HandleFunc("/", getDollarExchange)

	fmt.Println("Server started...")

	http.ListenAndServe(":8081", nil)
}

func getDollarExchange(writer http.ResponseWriter, request *http.Request) {

	dollarData := requestDollarExchange(writer)

	exchange := src.Exchange{}

	teste := src.Exchange{}
	err := json.NewDecoder(bytes.NewReader(dollarData)).Decode(&teste)
	if err != nil {
		panic(err)
	}

	fmt.Println(teste.USDBRL)

	err = json.Unmarshal(dollarData, &exchange)
	if err != nil {
		panic(err)
	}

	print(exchange.USDBRL.Name)

	storeExchange(dollarData)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	writer.Write([]byte(dollarData))
}

func storeExchange(dollarData []byte) {

}

func requestDollarExchange(writer http.ResponseWriter) []byte {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", dollarApiUrl, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// qual motivo fica aparecendo unhandled error e no vídeo realizando chamadas HTTP não?
	// posso por o defer antes da verificação de erro do http?
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func storeDollarExchange() {
	//db, err := gorm.Open(sqlite.Open("dollar-exchange.db"), &gorm.Config{})
	//if err != nil {
	//		panic(err)
	//	}

}
