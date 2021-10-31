package main

//slice boyutsuz bir dizi tanımında slice kullanılır bir boyur tanımına gerek yok array gibi
// array boyutlu limitli slice limitli değil

import "fmt"

var slc_1 []int

func main() {
	slc_2 := make([]int, 0, 3)

	//slc_2[0] = 2

	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 2)

	fmt.Println(slc_1, slc_2)

	// Append ile boyutu belli olmayan diziye boyutunu büyüterek atama yapılır kapasitenin 2 kartı kadar artar
	// Cap bir slice ne kadar kapasitesi var öğrenilebilir

	fmt.Println("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))

	fmt.Println("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))
}
