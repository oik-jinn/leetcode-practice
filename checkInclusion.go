package main

import "fmt"
import "os"

func main() {
	s1 := "rvwrk" 
	s2 := "lznomzggwrvrkxecjaq"

	fmt.Println(checkInclusion(s1, s2))
	os.Exit(1)
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := make(map[byte]int, 26)
	for i := 0; i < len(s1); i++ {
		_, ok := s1Map[s1[i]]
		if ok {
			s1Map[s1[i]]++
		} else {
			s1Map[s1[i]] = 1
		}
	}

	s2Map := make(map[byte]int, 26)
	for i := 0; i < len(s2); i++ {
		if i < len(s1) {
			s2Map[s2[i]]++
		} else {
			s2Map[s2[i]]++;
			s2Map[s2[i - len(s1)]]--;
		}

		if comapareMap(s1Map, s2Map) {
			return true
		}
	}

	return false
}

func comapareMap(s1Map map[byte]int, s2Map map[byte]int) bool {
	for k, v := range s1Map {
		if s2Map[k] != v {
			return false
		}
	}

	return true
}