package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writeFiles()

	readFiles()

	readPartFiles()

	removeFiles()
}

func writeFiles() {

	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!!!"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho: %d bytes \n", tamanho)

	defer f.Close()
}

func readFiles() {
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))
}

func readPartFiles() {
	arquivo, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}
}

func removeFiles() {
	err := os.Remove("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo removido com sucesso!!!")
}
