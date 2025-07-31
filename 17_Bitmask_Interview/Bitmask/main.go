package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Задание 1

func isSubmask(mask string, submask string) bool {
	maskB := ipToUint32(mask)
	submaskB := ipToUint32(submask)
	// ^maskB инвертирует маску, благодаря чему
	// последующее & приравняет все значения субмаски под изначальными единицами маски к 0,
	// а если у субмаски есть единицы под нулями у маски, & выдаст на их местах 1, из-за чего
	// число станет != 0
	return (submaskB & (^maskB)) == 0
}

func ipToUint32(ipStr string) uint32 {
	parts := strings.Split(ipStr, ".")
	// Преобразуем 4 числа и собираем их в одно большое, с которым будет удобно работать
	b1, _ := strconv.Atoi(parts[0])
	b2, _ := strconv.Atoi(parts[1])
	b3, _ := strconv.Atoi(parts[2])
	b4, _ := strconv.Atoi(parts[3])
	return (uint32(b1) << 24) | (uint32(b2) << 16) | (uint32(b3) << 8) | uint32(b4)
}

// Задание 2

// Сдвиг с поворотом влево
func rol(val uint32, shift uint) uint32 {
	//Делаем сдвиг на shift, а для выпавших битов делаем сдвиг
	// в другую сторону на 32 - shift, после чего объединяем |
	return (val << shift) | (val >> (32 - shift))
}

// Сдвиг с поворотом вправо
func ror(val uint32, shift uint) uint32 {
	return (val >> shift) | (val << (32 - shift))
}

func main() {

	//Задание 1

	Mask := "255.0.0.0"
	Submask := "127.0.0.1"
	fmt.Println(isSubmask(Mask, Submask)) // false
	Submask = "127.0.0.0"
	fmt.Println(isSubmask(Mask, Submask)) // true

	// Задание 2

	Val := uint32(0b11010000000000001000000000000010)
	fmt.Printf("%032b\n", rol(Val, 2)) //01000000000000100000000000001011
	fmt.Printf("%032b\n", ror(Val, 2)) //10110100000000000010000000000000

}
