package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {

	generateJson()

	readJson()
}

func generateJson() {
	println("========= Generate Json =========")

	//Marshal
	conta := Conta{Numero: 1, Saldo: 2}

	res, error := json.Marshal(conta)

	if error != nil {
		panic(error)
	}

	println(string(res))

	//Encoder
	encoder := json.NewEncoder(os.Stdout)

	encoder.Encode(conta)
}

func readJson() {
	println("========= Read Json =========")

	jsonText := []byte(`{"n":1,"s":2}`)

	var conta Conta

	err := json.Unmarshal(jsonText, &conta)

	if err != nil {
		panic(err)
	}

	println(conta.Numero)
}
