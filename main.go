package main

import (
	"Manabu/service"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetRandomFlashcard)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Listen and serve error:", err)
	}
}

func GetRandomFlashcard(w http.ResponseWriter, r *http.Request) {
	kanji := service.GetRandomKanji(1)

	word := service.GetRandomWordByKanji(kanji)

	fmt.Fprintf(w, fmt.Sprintf("Kanji: %s \nGlosses: %s \nPronounced %s \nWritten: %s",
		word.Kanji, word.Meanings[0].Glosses[0], word.Variants[0].Pronounced, word.Variants[0].Written))
}
