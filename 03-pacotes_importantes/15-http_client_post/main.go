package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}

	json := bytes.NewBuffer([]byte(`{"name": "gopher"}`))

	resp, err := c.Post("http://google.com", "application/json", json)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
