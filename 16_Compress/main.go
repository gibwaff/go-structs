package main

import "fmt"

//Задание 1
func compress(str string) string {
	var res, rework []byte

	for _, val := range str {
		rework = append(rework, byte(val))
	}

	cnt_num := 1
	for i := 0; i < len(rework); i++ {
		res = append(res, rework[i], 49)

		for i+1 < len(rework) {
			if rework[i+1] == rework[i] {
				res[cnt_num]++
				i++
			} else {
				break
			}
		}
		cnt_num += 2
	}

	return string(res)
}

//Задание 2
func decompress(str string) string {
	var res, rework []byte

	for _, val := range str {
		rework = append(rework, byte(val))
	}

	for i := 0; i < len(rework); i += 2 {
		for rework[i+1]-47 > 1 {
			res = append(res, rework[i])
			rework[i+1]--
		}
	}

	return string(res)
}

//Задание 3

func main() {
	//Задание 1
	work1 := "aaaabbbc1"
	fmt.Println(compress(work1)) // a4b3c111

	//Задание 2
	work2 := "ssssswwwwwaaaagaa"
	compressed_work2 := compress(work2)
	fmt.Println(compress(work2))              // s5w5a4g1a2
	fmt.Println(decompress(compressed_work2)) //ssssswwwwwaaaagaa

	//Задание 3

}
