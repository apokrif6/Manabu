package apiconfig

const kanjiGradeEndpoint = "https://kanjiapi.dev/v1/kanji/grade-"

const kanjiWordsEndpoint = "https://kanjiapi.dev/v1/words/"

func KanjiGradeEndpoint() string {
	return kanjiGradeEndpoint
}

func KanjiWordsEndpoint() string {
	return kanjiWordsEndpoint
}
