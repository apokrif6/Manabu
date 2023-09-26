package object

type Word struct {
	Kanji    string
	Meanings []struct {
		Glosses []string
	}
	Variants []struct {
		Priorities []interface{}
		Pronounced string
		Written    string
	}
}
