package viewmodels

type ProblemsVM struct {
	Problems []Problem `json:"problems"`
}

type Problem struct {
	ID      int    `json:"id"`
	Text    string `json:"problem_text"`
	Type    int    `json:"type"`
	Choices string `json:"choices"`
	Answer  string `json:"-"`
}
