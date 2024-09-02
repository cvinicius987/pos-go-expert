package main

type MyNumber int

type Number interface {
	~int | float64
}

func SomaInteiro(m map[string]int) int {
	soma := 0
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaConstraint[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	result := SomaInteiro(map[string]int{"a": 1, "b": 2})
	result2 := Soma(map[string]float64{"a": 1.5, "b": 2})
	result3 := SomaConstraint(map[string]int{"a": 1, "b": 10})
	result4 := SomaConstraint(map[string]MyNumber{"a": 20, "b": 10})

	println("Resultado SomaInteiro:", result)
	println("Resultado Soma:", result2)
	println("Resultado SomaConstraint:", result3)
	println("Resultado SomaConstraint + MyNumber:", result4)
}
