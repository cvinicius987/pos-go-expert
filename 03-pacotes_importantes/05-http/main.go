package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":3000", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := BuscaCep(cepParam)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*Cep, error) {

	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json")

	if error != nil {
		return nil, error
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		return nil, error
	}

	var cepResult Cep
	error = json.Unmarshal(body, &cepResult)

	if error != nil {
		return nil, error
	}

	return &cepResult, nil

}
