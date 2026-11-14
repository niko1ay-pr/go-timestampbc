package domain

type QuestionType string

const (
	QuestionTypeSingle   QuestionType = "single"
	QuestionTypeMultiple QuestionType = "multiple"
)

type Question struct {
	ID       string           `json:"id"`
	PollID   string           `json:"poll_id"`
	Title    string           `json:"title"`
	Type     QuestionType     `json:"type"`
	Sequence int              `json:"sequence"`
	Options  []QuestionOption `json:"question_options"`
}

type QuestionOption struct {
	ID         string `json:"id" db:"id"`
	QuestionID string `json:"question_id" db:"question_id"`
	Title      string `json:"title" db:"title"`
	Sequence   int    `json:"sequence" db:"sequence"`
}
