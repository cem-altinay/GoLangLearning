package main

import "fmt" // farklı bir paket kullanmak istenildiğinde  (.Net Using)

var arr_1 [3]int // default [0,0,0]
var arr_2 = [5]int{1, 2, 3, 4, 5}

func main() {
	// int array tanımlandı make ike tip ve kaç değer alacağı veriliyor
	arr_3 := make([]int, 3)

	arr_3[1] = 2
	fmt.Println(arr_1, arr_2, arr_3)

	fmt.Println("arr_1 len:%d \n", len(arr_1)) // len dizinin boyutunu verir
	fmt.Println("arr_2 len:%d \n", len(arr_2))
}
