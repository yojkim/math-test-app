package domains

type Problem struct {
	ID      int    `gorm:"AUTO_INCREMENT" json:"id"`
	Text    string `gorm:"problem_text" json:"problem_text"`
	Type    int    `gorm:"type" json:"type"`
	Choices string `gorm:"choices" json:"choices"`
	Answer  string `gorm:"answer" json:"answer"`
}
