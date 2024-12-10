package main

import "fmt"

type Player struct {
	Rating   int
	Nickname string
}

func searchRightmostPlayerWithRating(queue []Player, ratingBand int) int {
	var left, right = 0, len(queue) - 1
	var middle int
	for left != right {
		middle = ((left + right) + 1) / 2
		if queue[middle].Rating > ratingBand {
			right = middle - 1
		} else {
			left = middle
		}
	}
	if queue[right].Rating == ratingBand {
		return right
	} else {
		return -1
	}
}

func insertPlayerInQueueWithShift(queue []Player, index int, newPlayer Player) []Player {
	queue = append(queue, Player{})
	for i := len(queue) - 1; i > index; i-- {
		queue[i] = queue[i-1]
	}
	queue[index+1] = newPlayer
	return queue
}

func main() {
	var leaderBoard = []Player{
		{700, "LutyDemon"},
		{1000, "Sokol"},
		{1500, "Gribson"},
		{1500, "Bebroid"},
		{1700, "SigmaBoy"},
		{2000, "ChillGay"},
	}
	newPlayer := Player{1500, "Новичок"}
	post := searchRightmostPlayerWithRating(leaderBoard, newPlayer.Rating)
	if post != -1 {
		leaderBoard = insertPlayerInQueueWithShift(leaderBoard, post, newPlayer)
	}
	fmt.Println(leaderBoard)
}
