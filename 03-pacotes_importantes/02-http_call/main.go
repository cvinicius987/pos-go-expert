package main

import (
	"io"
	"net/http"
)

func main() {
	call, err := http.Get("https://www.uol.com.br")
	defer call.Body.Close()

	if err != nil {
		panic(err)
	}

	res, _ := io.ReadAll(call.Body)

	println(string(res))
}
