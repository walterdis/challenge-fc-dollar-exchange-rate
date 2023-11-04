package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/walterdis/fc-go-expert-challenges/src"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

const dollarApiUrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
const dbFile string = "database/exchange.sqlite"

func main() {
	http.HandleFunc("/cotacao", getDollarExchange)

	fmt.Println("Server started...")

	http.ListenAndServe(":8081", nil)
}

func getDollarExchange(writer http.ResponseWriter, request *http.Request) {

	dollarData := requestDollarExchange(writer)

	exchange := src.Exchange{}

	hydrateExchange(&exchange, dollarData)
	storeExchange(&exchange, writer)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(exchange.Bid)
}

func hydrateExchange(exchange *src.Exchange, dollarData []byte) {
	exchangeMap := map[string]map[string]any{}

	_ = json.Unmarshal(dollarData, &exchangeMap)

	jsonData, err := json.Marshal(exchangeMap["USDBRL"])

	err = json.Unmarshal(jsonData, &exchange)

	if err != nil {
		panic(err)
	}
}

func storeExchange(exchange *src.Exchange, writer http.ResponseWriter) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(src.Exchange{})

	if err := db.WithContext(ctx).Create(exchange).Error; err != nil {
		panic(err)
	}

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
