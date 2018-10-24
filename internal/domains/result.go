package domains

type Result struct {
	ID        int    `gorm:"AUTO_INCREMENT" json:"-"`
	ProblemID int    `gorm:"problem_id" json:"id"`
	Answer    string `gorm:"answer" json:"answer"`
	Result    int    `gorm:"result" json:"result"`
}
