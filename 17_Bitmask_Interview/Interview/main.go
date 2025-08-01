package main

import "fmt"

//Задание 1

//Первый вариант - плохой, т.к. работает за O(N^2)
func findSubarray1(arr []int, s int) bool {
	//создаём 2-мерный массив, в котором хранятся суммы
	// каждых элементов через "neighbour" элементов
	matrixArr := [][]int{}
	matrixArr = append(matrixArr, arr)
	line_cnt, neighbour := 0, 1 //считать, какую строку добавляем, какие отступы при сложении делать
	for i := 2; i < len(arr); i *= 2 {
		var newArr []int
		for k := 0; k < len(matrixArr[line_cnt])-neighbour; k++ {
			sum := matrixArr[line_cnt][k] + matrixArr[line_cnt][k+neighbour]
			newArr = append(newArr, sum)
			if sum == s {
				return true
			}
		}
		matrixArr = append(matrixArr, newArr)
		line_cnt++
		neighbour++
	}

	return false
}

//Второй вариант - работает за O(N), но не подходит для отрицательных чисел
func findSubarray2(arr []int, s int) bool {
	//искл. случай - 0
	if s == 0 {
		for _, i := range arr {
			if i == 0 {
				return true
			}
		}
		return false
	}

	left, sum := 0, 0
	for right := 0; right < len(arr); right++ {
		sum += arr[right]
		if sum > s {
			sum -= arr[left]
			left++
		}
		if sum == s {
			return true
		}
	}

	return false
}

//Третий вариант - универсальный за O(N), работает
// с помощью префиксных сумм и их записи в мапе
func findSubarray3(arr []int, s int) bool {
	prefixSum := 0
	seen := map[int]bool{}
	for _, el := range arr {
		prefixSum += el
		sum := prefixSum - s
		if seen[sum] {
			return true
		}
		seen[prefixSum] = true
	}

	return false
}

func main() {

	//Задание 1
	MyArray := []int{2, 3, 7, 6, 1, 3}
	S := 16
	fmt.Println(findSubarray3(MyArray, S)) //true

}
