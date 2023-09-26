package main

import (
	"Manabu/service"
	"fmt"
)

func main() {
	kanji := service.GetRandomKanji(1)

	word := service.GetRandomWordByKanji(kanji)

	fmt.Println("Kanji:", word.Kanji)
	fmt.Println("Glosses:", word.Meanings[0].Glosses[0])
	fmt.Println("Pronounced:", word.Variants[0].Pronounced)
	fmt.Println("Written:", word.Variants[0].Written)
}
