package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func sum(slice []int) int {
	var sum int
	for _, n := range slice {
		sum += n
	}
	return sum
}

func lightestTrack(track1, track2, track3 []int) int {
	if len(track1) == len(track2) && len(track1) == len(track3) {
		if len(track1) == 0 {
			return 1
		}
		if sum(track1) < sum(track2) {
			if sum(track1) < sum(track3) {
				return 1
			} else {
				return 3
			}
		} else {
			if sum(track2) < sum(track3) {
				return 2
			} else {
				return 3
			}
		}
	}

	if len(track1) < len(track2) {
		if len(track1) < len(track3) {
			return 1
		} else {
			return 3
		}
	} else {
		if len(track2) < len(track3) {
			return 2
		} else {
			return 3
		}
	}
}

func main() {
	flag.Parse()
	var hash = make(map[int]int)
	for _, i := range flag.Args() {
		appendHash(parseBoxInfo(i), hash)
	}
	track1 := []int{}
	track2 := []int{}
	track3 := []int{}
	for {
		if len(hash) == 0 {
			break
		}
		keyBiggest := findBiggest(hash)
		var lightest = lightestTrack(track1, track2, track3)
		switch lightest {
		case 1:
			track1 = append(track1, hash[keyBiggest])
		case 2:
			track2 = append(track2, hash[keyBiggest])
		case 3:
			track3 = append(track3, hash[keyBiggest])
		}
		delete(hash, keyBiggest)
	}
	fmt.Println("track1 = {}", track1)
	fmt.Println("track2 = {}", track2)
	fmt.Println("track3 = {}", track3)
	fmt.Println(sum(track1), sum(track2), sum(track3))
}

func findBiggest(hash map[int]int) int {
	var biggest = -1
	var keyBiggest = -1
	for key, value := range hash {
		if value > biggest {
			biggest = value
			keyBiggest = key
		}
	}
	//fmt.Println(biggest, keyBiggest)
	return keyBiggest
}

func appendHash(boxInfo []string, hash map[int]int) map[int]int {
	key, _ := strconv.Atoi(boxInfo[0])
	value, _ := strconv.Atoi(boxInfo[1])
	hash[key] = value
	return hash
}

func parseBoxInfo(argstr string) []string {
	return strings.Split(argstr, ":")
}
