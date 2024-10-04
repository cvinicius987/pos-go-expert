package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var viaCep = "https://viacep.com.br/ws/%s/json"

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {

		req, err := http.Get(fmt.Sprintf(viaCep, cep))

		if err != nil {
			panic(err)
		}

		defer req.Body.Close()

		resCepCall, err := io.ReadAll(req.Body)

		if err != nil {
			panic(err)
		}

		var resultCep Cep
		err = json.Unmarshal(resCepCall, &resultCep)

		file, err := os.Create(fmt.Sprintf("%s.json", cep))

		if err != nil {
			panic(err)
		}

		defer file.Close()

		formatFile := fmt.Sprintf("Endereço: %s", resultCep)
		_, err = file.WriteString(formatFile)

		fmt.Println(formatFile)
	}
}
