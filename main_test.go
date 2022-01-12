package main

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {
	data := []struct {
		Guess, Target string
		Result        string
	}{
		{"abcde", "fghij", "00000"},
		{"abcde", "abcde", "22222"},
		{"abcde", "abcgg", "22200"},
		{"abcde", "abced", "22211"},
		{"aaaab", "aggga", "21000"},
	}
	for _, tst := range data {
		t.Run(fmt.Sprintf("%s-%s", tst.Guess, tst.Target), func(t *testing.T) {
			result := score(tst.Guess, tst.Target)
			if result != tst.Result {
				t.Errorf("'%s' is not '%s'", result, tst.Result)
				t.Fail()
			}
		})
	}
}
