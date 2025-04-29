package main

import (
	"fmt"
	"math"
)

type TruckCoordinate struct {
	X int
	Y int
}

func CreateCoordinateHeap(a TruckCoordinate) []TruckCoordinate {
	return []TruckCoordinate{a}
}

func GetLength(coord TruckCoordinate) float64 {
	//высчитываем расстояние по Пифагору
	return math.Sqrt(math.Pow(float64(coord.X), 2) + math.Pow(float64(coord.Y), 2))
}

func AddCoordinates(heap []TruckCoordinate, coord TruckCoordinate) []TruckCoordinate {
	heap = append(heap, coord)
	for i := len(heap) - 1; GetLength(coord) < GetLength(heap[(i-1)/2]); i = (i - 1) / 2 {
		heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
	}
	return heap
}

func TruckCoordinatesHeap(coord []TruckCoordinate) []TruckCoordinate {
	heap := CreateCoordinateHeap(coord[0])
	for i := 1; i < len(coord); i++ {
		heap = AddCoordinates(heap, coord[i])
	}
	return heap
}

func GetPick(heap []TruckCoordinate) ([]TruckCoordinate, TruckCoordinate) {
	pick := heap[0]
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[:len(heap)-1]

	i := 0
	for {
		left := 2*i + 1
		right := 2*i + 2
		parent := i

		if right < len(heap) && GetLength(heap[right]) < GetLength(heap[parent]) {
			parent = right
		}
		if left < len(heap) && GetLength(heap[left]) < GetLength(heap[parent]) {
			parent = left
		}

		if parent == i {
			break
		}

		heap[i], heap[parent] = heap[parent], heap[i]
		i = parent
	}

	return heap, pick
}

// Конечная функция
func ClosestTrucks(heap []TruckCoordinate, k int) []TruckCoordinate {
	res_heap, pick := []TruckCoordinate{}, TruckCoordinate{}
	for i := 0; i < k; i++ {
		heap, pick = GetPick(heap)
		res_heap = append(res_heap, pick)
	}
	return res_heap
}

func main() {
	Trucks := []TruckCoordinate{{1, 2}, {-1, 3}, {4, 5}, {1, 0}}
	TrucksHeap := TruckCoordinatesHeap(Trucks)
	fmt.Println(ClosestTrucks(TrucksHeap, 2))
}
