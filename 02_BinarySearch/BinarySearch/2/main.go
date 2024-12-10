package main

import (
	"fmt"
	"math"
)

type DatingUser struct {
	iq   int
	name string
}

func sortUsersByIq(datingUsers []DatingUser) (usersSortedByIq []DatingUser) {
	for len(datingUsers) != 0 {
		var last_mini, index_mini = math.MaxInt32, -1
		for index, user := range datingUsers {
			if user.iq <= last_mini {
				last_mini = user.iq
				index_mini = index
			}
		}

		usersSortedByIq = append(usersSortedByIq, datingUsers[index_mini])
		datingUsers = removeElement(datingUsers, index_mini)
	}

	return
}

func removeElement(array []DatingUser, index int) []DatingUser {
	return append(array[:index], array[index+1:]...)
}

func findUsers(usersSortedByIq []DatingUser, lowerIqBound, professorIq int) (correctUsersName []string) {
	var lowerIqBound_index, professorIq_index int

	lowerIqBound_index = findUserIndex(usersSortedByIq, lowerIqBound)

	professorIq_index = findUserIndex(usersSortedByIq, professorIq+1)

	var correctUsers []DatingUser
	correctUsers = append(correctUsers, usersSortedByIq[lowerIqBound_index:professorIq_index]...)
	for _, el := range correctUsers {
		correctUsersName = append(correctUsersName, el.name)
	}

	return
}

func findUserIndex(userArray []DatingUser, iq int) (index int) {
	left, right := 0, len(userArray)-1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if userArray[middle].iq > iq {
			right = middle - 1
		} else if userArray[middle].iq < iq {
			left = middle + 1
		} else {
			index = middle
			break
		}
	}
	if index == 0 {
		index = (left+right)/2 + 1
	}

	return
}

func main() {
	var datingUsers = []DatingUser{
		{92, "Artem"},
		{86, "Daria"},
		{152, "IvanZolo"},
		{15, "Alex"},
		{32, "Anastsia"},
		{48, "Maria"},
		{80, "Nikita"},
		{-2, "Jack"},
	}
	ProfessorIq := 90
	LowerIqBound := 50

	usersSortedByIq := sortUsersByIq(datingUsers)
	correctUsers := findUsers(usersSortedByIq, LowerIqBound, ProfessorIq)
	fmt.Println(correctUsers)
}
