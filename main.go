package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"unicode"
)

type L struct {
	Word []byte
	cnt  int
}

func main() {

	l := &[1]L{}
	l2 := l[:]

	//check cnt args
	if len(os.Args) < 2 {
		log.Println("Missing parameter, provide file name!")
		return
	}

	//get data
	tmp := readFile(os.Args[1])

	// get count words and theirs position
	cntWords, _, data := wordCount(tmp)

	//count words
	for i := 0; i < cntWords; i++ {

		k := check(data[i][:], l2)
		if k > -1 {
			l2[k].cnt++
		} else {
			tmp2 := Newl(data[i][:])

			l2 = append(l2, *tmp2)
		}
	}

	//sort words by cnt
	sort.Slice(l2, func(i, j int) (less bool) {
		return l2[i].cnt > l2[j].cnt
	})
	for i := 0; i < 20; i++ {
		fmt.Printf("%d %s\n", l2[i].cnt, l2[i].Word)
	}
}

func checktwoWords(data []byte) int {

	//ln := len(data)

	//i := bytes.IndexByte(data, ' ')
	if ok := bytes.Contains(data, []byte(" ")); ok {
		for i := range data {
			if i > 0 && unicode.IsSpace(rune(data[i])) {
				tmp := data[i+1:]
				if len(tmp) > 1 {
					return i + 1
				}
			}
		}
	}

	return -1
}

//wordCount count words
func wordCount(value []byte) (int, [][]int, [][]byte) {
	// Match only ads
	re := regexp.MustCompile(`[A-Za-z]+`)

	// Find all matches and return count.
	res := re.FindAll(value, -1)

	//replase char, dis and do lowercase
	for i := range res {
		res[i] = bytes.ToLower(res[i])
	}
	inx := re.FindAllIndex(value, len(res))

	for k := range res {
		i := checktwoWords(res[k])
		if i > 0 {
			tmp := re.FindAll(res[k], -1)

			res = append(res, tmp...)
		}
	}

	return len(res), inx, res
}

func Newl(word []byte) *L {
	return &L{
		Word: word,
		cnt:  1,
	}
}

//check if word in mass L than return index word
func check(word []byte, words []L) int {

	for i := 0; i < len(words); i++ {
		flag := bytes.Equal(word, words[i].Word)
		if flag {
			return i
		}
	}
	return -1
}

// r Read data from file or err Fatal]
func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("can't read file: ", os.Args[1])
		log.Fatal(err)
	}
	return data
}
