package main

import "fmt"

// map: key value değerler tutar sıra garantisi olmadan adreste tutat

var map_1 map[int]string

func main() {
	map_1 = make(map[int]string) // map_1 initilize

	map_2 := make(map[int]int)
	map_3 := make(map[int]int, 3)

	map_1[0] = "1"
	map_2[0] = 2
	map_3[0] = 3

	map_3[1] = 3
	map_3[2] = 3
	map_3[3] = 3

	fmt.Println(map_1, map_2, map_3)
	fmt.Println("map_1 len: \n", len(map_1))
	fmt.Println("map_2len: \n", len(map_2))
	fmt.Println("map_3 len: \n", len(map_3))
}
