package generalG

import (
	"fmt"
	"strings"
)

var (
	num           int
	roman         string
	rep           int = 0
	mappedRoman   int
	nextChar      int
	err           error
	numToRomanMap map[int]string = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	romanToNumMap map[string]int = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

// arguments:
// s = roman numeral
// returns int
// @TODO invalid multiple small value char on the left of big value char ex. IIIV,
// incorect rep count for char I X C M
func RomanToNum(input string) (int, error) {
	s := strings.ToUpper(input)
	l := len(s)

	for i := 0; i < l; i++ {
		roman = string(s[i])
		mappedRoman = romanToNumMap[roman]
		if isEmpty(mappedRoman) {
			return 0, fmt.Errorf("invalid roman character : %s", string(input[i]))
		}

		if i < l-1 {
			// these roman chars can only repeated 3 times.
			if s[i] == s[i+1] { // compare to next index
				rep, err = RepetationConstraint(rep, roman)
				if err != nil {
					return 0, err
				}
			}
			if nextChar = romanToNumMap[string(s[i+1])]; mappedRoman < nextChar {
				num += nextChar - mappedRoman
				i++
			} else {
				num += mappedRoman
			}
		} else {
			// these roman chars can only repeated 3 times.
			rep, err = RepetationConstraint(rep, roman)
			if err != nil {
				return 0, err
			}
			num += mappedRoman
		}
	}

	return num, nil
}

func RepetationConstraint(rep int, char string) (int, error) {
	if char == "I" || char == "X" || char == "C" || char == "M" {
		rep++
		if rep > 3 {
			return 0, fmt.Errorf("repetition more than 3 times of either \"I\", \"X\", \"C\", or \"M\" character is invalid")
		}
	} else {
		rep = 0
	}
	return rep, nil
}
