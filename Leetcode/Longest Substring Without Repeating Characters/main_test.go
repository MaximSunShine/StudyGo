package main

import (
	"testing"
)

type solution struct {
	text   string
	answer int
}

func TestLengthOfLongestSubstring(t *testing.T) {

	s := []solution{
		{"", 0},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, v := range s {
		if rez := lengthOfLongestSubstring(v.text); rez != v.answer {
			t.Errorf("get '%s' Expected '%d', but got '%d'", v.text, v.answer, rez)
		}
	}

}
