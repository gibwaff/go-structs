package main

import "fmt"

type Warehouse struct {
	FreeAt int
	Index  int
}

func CreateHeap(a Warehouse) []Warehouse {
	return []Warehouse{a}
}

func AddToHeap(heap []Warehouse, warehouse Warehouse) []Warehouse {
	heap = append(heap, warehouse)
	for i := len(heap) - 1; warehouse.FreeAt < heap[(i-1)/2].FreeAt; i = (i - 1) / 2 {
		heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
	}
	return heap
}

func ArrayToHeap(arr []Warehouse) []Warehouse {
	NewHeap := CreateHeap(arr[0])
	for i := 1; i < len(arr); i++ {
		NewHeap = AddToHeap(NewHeap, arr[i])
	}
	return NewHeap
}

func GetPick(heap []Warehouse) ([]Warehouse, Warehouse) {
	pick := heap[0]
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[:len(heap)-1]

	i := 0
	for {
		left := i*2 + 1
		right := i*2 + 2
		parent := i

		if right < len(heap) && heap[right].FreeAt < heap[parent].FreeAt {
			parent = right
		}
		if left < len(heap) && heap[left].FreeAt < heap[parent].FreeAt {
			parent = left
		}

		if i == parent {
			break
		}

		heap[i], heap[parent] = heap[parent], heap[i]
		i = parent

	}

	return heap, pick
}

func UnloadingTruck(n int, times []int) []int {
	Warehouses := []Warehouse{}
	for i := 1; i <= n; i++ {
		Warehouses = append(Warehouses, Warehouse{0, i})
	}
	WarehousesHeap := ArrayToHeap(Warehouses)

	ResArray := make([]int, len(times))
	for i := 0; i < len(times); i++ {
		UsingWH := Warehouse{}
		WarehousesHeap, UsingWH = GetPick(WarehousesHeap)
		ResArray[i] = UsingWH.FreeAt
		UsingWH.FreeAt += times[i]
		WarehousesHeap = AddToHeap(WarehousesHeap, UsingWH)
	}

	return ResArray
}

func main() {
	Trucks := []int{5, 8, 3, 4, 1, 10, 2}
	fmt.Println(UnloadingTruck(3, Trucks))
}
