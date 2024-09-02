package main

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func main() {
	println(Compara(1, 1))
}
