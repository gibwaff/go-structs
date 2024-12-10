package main

import "fmt"

func main() {
	var newRussianDictionary = []string{
		"Контент",
		"Лутер",
		"Тренд",
		"Фиксер",
		"Фэшн",
		"Хипстер",
	}

	insertNewWord := func(word string) {
		runeWord := []rune(word)
		var left, right = 0, len(newRussianDictionary)
		var middle int
		for left < right {
			middle = (left + right) / 2
			if []rune(newRussianDictionary[middle])[0] < runeWord[0] {
				left = middle + 1
			} else {
				right = middle
			}
		}
		if left == len(newRussianDictionary) {
			newRussianDictionary = append(newRussianDictionary, word)
			return
		}
		newRussianDictionary = append(newRussianDictionary, "")
		fl := true
		for fl {
			len_fl := true
			runeWordInDict := []rune(newRussianDictionary[left])
			lower_length := min(len(runeWordInDict), len(runeWord))
			for i := 0; i < lower_length; i++ {
				if runeWord[i] != runeWordInDict[i] {
					len_fl = false
					if runeWordInDict[i] > runeWordInDict[i] {
						left++
						break
					} else {
						fl = false
						break
					}
				}
			}
			if len_fl {
				if len(runeWordInDict) < len(runeWord) {
					left++
				} else {
					fl = false
				}
			}
		}
		for i := len(newRussianDictionary) - 1; i > left; i-- {
			newRussianDictionary[i] = newRussianDictionary[i-1]
		}
		newRussianDictionary[left] = word
	}

	insertNewWord("Бебра")
	insertNewWord("Свэг")
	insertNewWord("Токсик")
	insertNewWord("Контентмейкер")
	insertNewWord("Абаюдна")
	insertNewWord("Хип")
	insertNewWord("Ямейн")

	fmt.Println(newRussianDictionary)
}
