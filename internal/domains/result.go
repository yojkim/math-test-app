package domains

type Answer struct {
	ProblemID int    `json:"id"`
	Answer    string `json:"answer"`
}

type Result struct {
	ID        int    `gorm:"AUTO_INCREMENT" json:"-"`
	ProblemID int    `gorm:"problem_id" json:"id"`
	Answer    string `gorm:"answer" json:"answer"`
	Result    bool   `gorm:"result" json:"result"`
}
