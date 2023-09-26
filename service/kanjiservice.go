package service

import (
	"Manabu/config"
	"Manabu/object"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func GetKanjiList(grade int) []string {
	resp, err := http.Get(apiconfig.KanjiGradeEndpoint() + strconv.Itoa(grade))

	if err != nil {
		return []string{}
	}

	body, _ := io.ReadAll(resp.Body)

	dataJson := string(body)

	var arr []string

	_ = json.Unmarshal([]byte(dataJson), &arr)

	return arr
}

func GetRandomKanji(grade int) string {
	kanjiList := GetKanjiList(grade)

	return kanjiList[rand.Intn(len(kanjiList)-1)]
}

func GetWordsByKanji(kanji string) []object.Word {
	if kanji == "" {
		return []object.Word{}
	}

	resp, err := http.Get(apiconfig.KanjiWordsEndpoint() + kanji)

	if err != nil {
		return []object.Word{}
	}

	body, _ := io.ReadAll(resp.Body)

	dataJson := string(body)

	var words []object.Word

	if err := json.Unmarshal([]byte(dataJson), &words); err != nil {
		fmt.Println("Can't unmarshall words by kanji:", err)
		return []object.Word{}
	}

	return words
}

func GetRandomWordByKanji(kanji string) object.Word {
	words := GetWordsByKanji(kanji)

	if len(words) == 0 {
		return object.Word{}
	}

	word := words[rand.Intn(len(words)-1)]

	word.Kanji = kanji

	return word
}
