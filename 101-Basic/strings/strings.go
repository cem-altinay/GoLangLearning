package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "lorem ipsum dolar adadsas fadasd dadad"
	str_1 := str[:5]
	str_2 := str[len(str)-4:]
	str_3 := fmt.Sprintf("%s ipsum dolar ad %s", str_1, str_2)

	if strings.EqualFold(str_1, "LOrem") {
		fmt.Println("str_1 equal to LOrem")
	}

	//https://www.youtube.com/watch?v=wAUL4f0F8hA

	fmt.Println(str)
	fmt.Println(str_1)
	fmt.Println(str_2)
	fmt.Println(str_3)
	fmt.Println(strings.ToUpper(str))

}
