package domain

type QuizModel struct {
	Id      string   `json:"id"`
	Soal    string   `json:"soal"`
	Jawaban []string `json:"jawaban"`
	Benar   string   `json:"jawabanBenar"`
	Score   int      `json:"score"`
	Tips    string   `json:"tips"`
}
