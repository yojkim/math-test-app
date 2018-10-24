package viewmodels

type ResultsVM struct {
	Results []Result `json:"results"`
}

type Result struct {
	ID        int    `json:"-"`
	ProblemID int    `json:"id"`
	Answer    string `json:"answer"`
	Result    bool   `json:"result"`
}
