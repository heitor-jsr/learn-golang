package main

import "fmt"

func main() {
	var varivael1 string = "var 1"
	var2 := "var2"
	fmt.Println(varivael1, var2)

	var (
		var3 string = "123321"
		var4 string = "fwewe"
	)

	fmt.Println(var3, var4)

	var5, var6 := "211", "424"

	fmt.Println(var5, var6)
	//inversão de valores de variaveis
	var5, var6 = var6, var5
	fmt.Println(var5, var6)

}
