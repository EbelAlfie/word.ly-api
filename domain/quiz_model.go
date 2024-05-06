package domain

type QuizModel struct (
	soal string `json:"soal"`
	nilai string `json:"nilai"`
	jawaban string `json:"jawaban"`
)

type PilihanGanda struct (
	jawaban string `json:"jawaban"`
	skor bool `json:"skor"`
)
