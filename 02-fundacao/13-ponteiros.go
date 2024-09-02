package main

func main() {
	texto := "Teste 1"

	println("Texto original:", texto)

	changeText(&texto)

	println("Texto alterado:", texto)
}

func changeText(txt *string) {
	*txt = "Teste 2"
	println("Alterando para: ", *txt)
}
