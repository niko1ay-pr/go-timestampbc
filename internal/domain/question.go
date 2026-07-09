package domain

type QuestionType string

const (
	QuestionTypeSingle   QuestionType = "single"
	QuestionTypeMultiple QuestionType = "multiple"
)

type Question struct {
	ID       string       `json:"id" db:"id"`
	PollID   string       `json:"poll_id" db:"poll_id"`
	Title    string       `json:"title" db:"title"`
	Type     QuestionType `json:"type" db:"type"`
	Sequence int          `json:"sequence" db:"sequence"`
}

type QuestionOption struct {
	ID         string `json:"id" db:"id"`
	QuestionID string `json:"question_id" db:"question_id"`
	Title      string `json:"title" db:"title"`
	Sequence   int    `json:"sequence" db:"sequence"`
}
