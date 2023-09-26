package main

import (
	"Manabu/service"
	"fmt"
)

func main() {
	kanji := service.GetRandomKanji(1)

	fmt.Printf(kanji)
}
