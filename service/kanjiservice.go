package service

import (
	"Manabu/config"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func GetKanjiList(grade int) []string {
	resp, _ := http.Get(apiconfig.KanjiGradeEndpoint() + strconv.Itoa(grade))

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
