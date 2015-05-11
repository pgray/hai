package main

import (
	"fmt"
	//	"io/ioutil"
	"net/http"
	//	"strings"
	//	"appengine"
	//  "appengine/urlfetch"
)

/*Each letter is 5 units high and 3 units wide.
BubbleLetter stores 5 numbers between 0 and 7,
that, when converted to binary, represent the
existence or absence of a face in each unit.
example "e","E":
7: 1 1 1
4: 1 0 0
7: 1 1 1
4: 1 0 0
7: 1 1 1
*/
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
		return BubbleLetter{First: 2, Second: 5, Third: 7, Fourth: 5, Fifth: 5}
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
		return BubbleLetter{First: 5, Second: 7, Third: 5, Fourth: 5, Fifth: 5}
	case "n", "N":
		return BubbleLetter{First: 5, Second: 7, Third: 7, Fourth: 7, Fifth: 5}
	case "o", "O":
		return BubbleLetter{First: 7, Second: 5, Third: 5, Fourth: 5, Fifth: 7}
	case "p", "P":
		return BubbleLetter{First: 7, Second: 5, Third: 7, Fourth: 4, Fifth: 4}
	case "q", "Q":
		return BubbleLetter{First: 7, Second: 5, Third: 5, Fourth: 7, Fifth: 1}
	case "r", "R":
		return BubbleLetter{First: 7, Second: 5, Third: 6, Fourth: 5, Fifth: 5}
	case "s", "S":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 1, Fifth: 7}
	case "t", "T":
		return BubbleLetter{First: 7, Second: 2, Third: 2, Fourth: 2, Fifth: 2}
	case "u", "U":
		return BubbleLetter{First: 5, Second: 5, Third: 5, Fourth: 5, Fifth: 7}
	case "v", "V":
		return BubbleLetter{First: 5, Second: 5, Third: 5, Fourth: 5, Fifth: 2}
	case "w", "W":
		return BubbleLetter{First: 7, Second: 4, Third: 7, Fourth: 4, Fifth: 7}
	case "x", "X":
		return BubbleLetter{First: 5, Second: 5, Third: 2, Fourth: 5, Fifth: 5}
	case "y", "Y":
		return BubbleLetter{First: 5, Second: 5, Third: 2, Fourth: 2, Fifth: 2}
	case "z", "Z":
		return BubbleLetter{First: 7, Second: 1, Third: 2, Fourth: 4, Fifth: 7}
	default:
		return BubbleLetter{First: 5, Second: 0, Third: 2, Fourth: 0, Fifth: 7}
	}
}

func bubbleWordField(word string) []BubbleLetter {
	wordField := make([]BubbleLetter, len(word))
	for i, _ := range wordField {
		wordField[i] = bubbleLetterField(string(word[i]))
	}
	return wordField
}

func fieldToString(i int) string {
	switch i {
	case 0:
		return "                        "
	case 1:
		return "                :simple_smile:"
	case 2:
		return "        :simple_smile:        "
	case 3:
		return "        :simple_smile::simple_smile:"
	case 4:
		return ":simple_smile:                "
	case 5:
		return ":simple_smile:        :simple_smile:"
	case 6:
		return ":simple_smile::simple_smile:        "
	case 7:
		return ":simple_smile::simple_smile::simple_smile:"
	default:
		return "Somehow, the integer passed to fieldToString was not 0-7..."
	}
}

func wordFieldToString(wordField []BubbleLetter) string {
	var firstString []byte
	var secondString []byte
	var thirdString []byte
	var fourthString []byte
	var fifthString []byte

	for i, _ := range wordField {
		firstString = append(firstString, fieldToString(wordField[i].First)...)
		secondString = append(secondString, fieldToString(wordField[i].Second)...)
		thirdString = append(thirdString, fieldToString(wordField[i].Third)...)
		fourthString = append(fourthString, fieldToString(wordField[i].Fourth)...)
		fifthString = append(fifthString, fieldToString(wordField[i].Fifth)...)
		//add space after each letter
		firstString = append(firstString, "    "...)
		secondString = append(secondString, "    "...)
		thirdString = append(thirdString, "    "...)
		fourthString = append(fourthString, "    "...)
		fifthString = append(fifthString, "    "...)
	}
	firstString = append(firstString, "\n"...)
	secondString = append(secondString, "\n"...)
	thirdString = append(thirdString, "\n"...)
	fourthString = append(fourthString, "\n"...)
	responseString := append(firstString, secondString...)
	responseString = append(responseString, thirdString...)
	responseString = append(responseString, fourthString...)
	responseString = append(responseString, fifthString...)
	return string(responseString)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("text"))

	thing := bubbleWordField(r.FormValue("text"))

	fmt.Println(thing[0].First, thing[0].Second, thing[0].Third, thing[0].Fourth, thing[0].Fifth)

	fmt.Fprintf(w, "%s", wordFieldToString(bubbleWordField(r.FormValue("text"))))

	fmt.Println("After write")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
