package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var monsters []string
	monsters = append(monsters, args...)
	insertionSort(monsters)
}

func getBattleResult(monster1, monster2 string) (string, string) {
	url := fmt.Sprintf("https://ob6la3c120.execute-api.ap-northeast-1.amazonaws.com/Prod/battle/%x+%x", monster1, monster2)
	response, err := http.Get(url)
	if err != nil {
		panic("Get request failed")
	}

	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)
	splittedResult := strings.Split(string(byteArray), "\"")
	winner, _ := hex.DecodeString(splittedResult[3])
	loser, _ := hex.DecodeString(splittedResult[7])
	return string(winner), string(loser)
}

func insertionSort(monsters []string) {
	//partition := 0
	frontline := 1
	current := 1
	for {
		if frontline >= len(monsters) {
			break
		}
		//fmt.Println(monsters[current-1], " VS ", monsters[current])
		winner, _ := getBattleResult(monsters[current-1], monsters[current])
		//fmt.Println("winner = ", winner)

		if winner != monsters[current] {
			frontline += 1
			current = frontline
			continue
		}

		if winner != monsters[current-1] {
			swap(monsters, current-1, current)
			//fmt.Println("swapped to ", monsters)
			current -= 1
		}

		if current == 0 {
			frontline += 1
			current = frontline
			continue
		}
	}

	fmt.Println(monsters)
}

func swap(arr []string, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
