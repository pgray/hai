package main

import (
	"fmt"
	//	"io/ioutil"
	"net/http"
	//	"strings"

	//	"appengine"
	//  "appengine/urlfetch"
)

type BubbleLetter struct {
	First  int
	Second int
	Third  int
	Fourth int
	Fifth  int
}

func bubbleLetterField(letter string) BubbleLetter {
	fmt.Println("LETTER:", letter)
	switch letter {
	case "a", "A":
		return BubbleLetter{First: 7, Second: 5, Third: 7, Fourth: 5, Fifth: 5}
	case "b", "B":
		return BubbleLetter{First: 7, Second: 5, Third: 6, Fourth: 5, Fifth: 7}
	case "c", "C":
		return BubbleLetter{First: 7, Second: 4, Third: 4, Fourth: 4, Fifth: 7}
	case "d", "D":
		return BubbleLetter{First: 6, Second: 5, Third: 5, Fourth: 5, Fifth: 6}
	case "e", "E":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "f", "F":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 4}
	case "g", "G":
		return BubbleLetter{First: 7, Second: 4, Third: 5, Fourth: 5, Fifth: 7}
	case "h", "H":
		return BubbleLetter{First: 5, Second: 5, Third: 7, Fourth: 5, Fifth: 5}
	case "i", "I":
		return BubbleLetter{First: 7, Second: 2, Third: 2, Fourth: 2, Fifth: 7}
	case "j", "J":
		return BubbleLetter{First: 1, Second: 1, Third: 1, Fourth: 5, Fifth: 7}
	case "k", "K":
		return BubbleLetter{First: 5, Second: 6, Third: 4, Fourth: 6, Fifth: 5}
	case "l", "L":
		return BubbleLetter{First: 4, Second: 4, Third: 4, Fourth: 4, Fifth: 7}
	case "m", "M":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "n", "N":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "o", "O":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "p", "P":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "q", "Q":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "r", "R":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "s", "S":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "t", "T":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "u", "U":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "v", "V":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "w", "W":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "x", "X":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "y", "Y":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "z", "Z":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	}
}

func bubbleWordField(word string) []BubbleLetter {
	wordField := make([]BubbleLetter, len(word))
	for i, _ := range wordField {
		wordField[i] = bubbleLetterField(string(word[i]))
	}
	return wordField
}

func handler(w http.ResponseWriter, r *http.Request) {
	thing := bubbleWordField("Word")
	fmt.Fprintf(w, "%s %d %d %d %d %d", r.FormValue("text"), thing[0].First, thing[0].Second, thing[0].Third, thing[0].Fourth, thing[0].Fifth)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
