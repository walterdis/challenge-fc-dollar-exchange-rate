package main

import (
	"encoding/json"
	"fmt"
	"github.com/walterdis/fc-go-expert-challenges/multithreadring/Entity"
	"io"
	"net/http"
	"time"
)

const (
	apiCepUrl string = "https://cdn.apicep.com/file/apicep/%s.json"
	viaCepUrl string = "https://viacep.com.br/ws/%s/json/"
)

func main() {
	chan1 := make(chan []byte)
	chan2 := make(chan []byte)

	go func() {
		cep, err := requestCep(apiCepUrl, "88075-120")
		if err != nil {
			panic(err)
		}
		chan1 <- cep
	}()

	go func() {
		cep, err := requestCep(viaCepUrl, "88040160")
		if err != nil {
			panic(err)
		}
		chan2 <- cep
	}()

	select {
	case cep1 := <-chan1:
		printFromApiCep(cep1)

	case cep2 := <-chan2:
		printFromViaCep(cep2)

	case <-time.After(time.Second * 1):
		println("Timeout")
	}
}

func requestCep(urlCep string, cep string) ([]byte, error) {
	url := fmt.Sprintf(urlCep, cep)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func printFromViaCep(cep []byte) {
	viaCep := Entity.ViaCep{}

	json.Unmarshal(cep, &viaCep)

	fmt.Println("API: ViaCep" +
		"\nCep: " + viaCep.Cep +
		"\nLogradouro: " + viaCep.Logradouro +
		"\nComplemento: " + viaCep.Complemento +
		"\nBairro: " + viaCep.Bairro +
		"\nLocalidade: " + viaCep.Localidade +
		"\nUf: " + viaCep.Uf +
		"\nIbge: " + viaCep.Ibge +
		"\nGia: " + viaCep.Gia +
		"\nDdd: " + viaCep.Ddd +
		"\nSiafi: " + viaCep.Siafi)
}

func printFromApiCep(cep []byte) {
	apiCep := Entity.ApiCep{}
	json.Unmarshal(cep, &apiCep)

	fmt.Println("API: ApiCep" +
		"\nCep: " + apiCep.Cep +
		"\nLogradouro: " + apiCep.Logradouro +
		"\nComplemento: " + apiCep.Complemento +
		"\nBairro: " + apiCep.Bairro +
		"\nLocalidade: " + apiCep.Localidade +
		"\nUf: " + apiCep.Uf +
		"\nIbge: " + apiCep.Ibge +
		"\nGia: " + apiCep.Gia +
		"\nDdd: " + apiCep.Ddd +
		"\nSiafi: " + apiCep.Siafi)

}
