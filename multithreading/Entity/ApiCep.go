package Entity

type ApiCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"address"`
	Complemento string `json:"complement"`
	Bairro      string `json:"district"`
	Localidade  string `json:"city"`
	Uf          string `json:"state"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}
