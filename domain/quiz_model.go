package domain

import "strconv"

type QuizRequest struct {
	QuizId        string   `json:"quizId"`
	Question      string   `json:"question"`
	ChoiceId      string   `json:"choiceId"`
	Choices       []string `json:"choices"`
	CorrectAnswer string   `json:"correctAnswer"`
	Score         int      `json:"score"`
	Hint          string   `json:"hint"`
	Type          int      `json:"type"`
}

type QuizModel struct {
	Id            string   `json:"id"`
	TId           string   `json:"teacherId"`
	Question      string   `json:"question"`
	ChoiceId      string   `json:"choiceId"`
	Choices       []string `json:"choices"`
	CorrectAnswer string   `json:"correctAnswer"`
	Score         int      `json:"score"`
	Hint          string   `json:"hint"`
	Type          int      `json:"type"`
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
