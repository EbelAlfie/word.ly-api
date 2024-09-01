package domain

import "strconv"

type QuizModel struct {
	Id      string   `json:"id"`
	TId     string   `json:"teacherId"`
	Soal    string   `json:"soal"`
	Jawaban []string `json:"jawaban"`
	Benar   string   `json:"jawabanBenar"`
	Score   int      `json:"score"`
	Tips    string   `json:"tips"`
}

type QuizType int

const (
	ALL             QuizType = 0
	CERPEN          QuizType = 1
	KALIMAT_EFEKTIF QuizType = 2
)

func ParseToEnum(rawValue string) QuizType {
	quizType, err := strconv.Atoi(rawValue)
	if err != nil {
		return ALL
	}
	switch quizType {
	case int(CERPEN):
		return CERPEN
	case int(KALIMAT_EFEKTIF):
		return KALIMAT_EFEKTIF
	default:
		return ALL
	}
}
