package main

import "fmt"

func main() {
	var value int32 = 9

	fmt.Println("address memory is", &value, "and value:", value)
	printAsCopy(value)
	fmt.Println("After function address memory is", &value, "and value:", value)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("address memory is", &value, "and value:", value)
	printAsReference(&value)
	fmt.Println("After function address memory is", &value, "and value:", value)
}

func printAsReference(value *int32) {
	fmt.Println("printAsReference -> address memory is", value, "and value:", *value)
	*value = 16 // Alterando o valor na memória, teremos efeito diretamente na função main
}

func printAsCopy(value int32) {
	fmt.Println("printAsCopy -> address memory is", &value, "and value:", value)
	value = 16 // como estamos mudando o valor somente aqui na variável de cópia, não teremos
	// o efeito na função main
}
