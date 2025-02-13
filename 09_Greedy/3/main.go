package main

import (
	"fmt"
	"strconv"
)

// Сортировка по стоимости товара за единицу объёма
func sortProds(arr [][]int) [][]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j][1]/arr[j][0] < arr[j+1][1]/arr[j+1][0] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	capacity := 85
	products := [][]int{
		{10, 550}, {5, 250}, {20, 1200}, {8, 400}, {15, 900},
		{12, 600}, {7, 350}, {18, 1100}, {6, 300}, {14, 870},
		{9, 450}, {16, 950}, {11, 550}, {13, 701}, {19, 1150},
	}
	sorted_products := sortProds(products)
	// sorted [[14 870] [18 1100] [20 1200] [15 900]
	// [19 1150] [16 950] [10 550] [13 701] [5 250]
	// [8 400] [12 600] [7 350] [6 300] [9 450] [11 550]]

	total_price := 0
	placed := 0
	el := 0
	max_price := 0

	//Загружаем, начиная с самых дорогих
	for placed+sorted_products[el][0] <= capacity {
		if max_price < sorted_products[el][1] {
			max_price = sorted_products[el][1]
		}
		placed += sorted_products[el][0]
		total_price += sorted_products[el][1]
		el++
	}

	//Оставшееся место забиваем частью невошедшего товара
	empty := capacity - placed
	total_price += sorted_products[el][1] * empty / sorted_products[el][0]

	fmt.Println("Цена за самый дорогой товар: " + strconv.Itoa(sorted_products[0][1]))
	fmt.Println("Цена за самый прибыльный товар: " + strconv.Itoa(max_price))

}
