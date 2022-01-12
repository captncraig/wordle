package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

var wordList = loadWords()

func findbest(list []string) string {
	portion := 1 / float64(len(wordList))
	min := float64(99999)
	minAt := ""
	for _, p := range list {
		lookups := map[string]int{}
		for _, p2 := range list {
			s := score(p, p2)
			lookups[s] = lookups[s] + 1
		}
		expected := float64(0)
		for _, count := range lookups {
			expected += float64(count) * (float64(count) * portion)
		}
		if expected < min {
			log.Println(p, expected)
			min = expected
			minAt = p
		}
	}
	return minAt
}

func main() {
	possible := wordList
	for {
		log.Printf("%d possible words:", len(possible))
		myGuess := findbest(possible)
		log.Println(myGuess)

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		newPossible := make([]string, 0, len(possible))
		for _, p := range possible {
			if score(myGuess, p) == text {
				newPossible = append(newPossible, p)
			}
		}
		possible = newPossible
	}
}

func loadWords() []string {
	dat, err := os.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(dat), "\n")
	sort.Strings(s)
	return s
}

func score(guess, target string) string {
	if len(guess) != len(target) {
		return "00000"
	}
	g := []byte(guess)
	t := []byte(target)
	result := []byte("00000")

	// first find direct hits
	for i, c := range t {
		if g[i] == c {
			result[i] = '2'
			t[i] = 0
		}
	}
	for i := range result {
		if result[i] == '2' {
			continue
		}
		gc := g[i]
		for j := range t {
			if result[j] != 0 && t[j] == gc {
				t[j] = 0
				result[i] = '1'
			}
		}
	}

	return string(result)
}
