package main

import "fmt"

func Pow(num, power int) (res int) {
	res = num
	if power == 0 {
		return 1
	}
	for i := 1; i < power; i++ {
		res *= num
	}
	return res
}

func RecursionCheck(cur_index, cur_sum, num int, PowersOfThree []int, res int) int {
	if cur_sum == num {
		return 1
	}
	for ; cur_index < len(PowersOfThree); cur_index++ {
		new_sum := cur_sum + PowersOfThree[cur_index]
		res += RecursionCheck(cur_index+1, new_sum, num, PowersOfThree, res)
	}
	return res
}

func CheckPowersOfThree(num int) bool {
	PowersOfThree := []int{}
	for cur := 0; Pow(3, cur) <= num; cur++ {
		PowersOfThree = append(PowersOfThree, Pow(3, cur))
	}
	if RecursionCheck(0, 0, num, PowersOfThree, 0) > 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(CheckPowersOfThree(91))
}
