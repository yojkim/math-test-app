package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yojkim/math-test-app/internal/domains"
)

type ResultRepository struct {
	Conn *gorm.DB
}

func (repo *ResultRepository) Store(result domains.Result) (int, error) {
	err := repo.Conn.Create(&result).Error
	if err != nil {
		return -1, err
	}

	return result.ID, nil
}

func (repo *ResultRepository) Check(answer *domains.Answer) (*domains.Result, error) {
	p := &domains.Problem{}
	res := &domains.Result{}

	err := repo.Conn.First(p, answer.ProblemID).Error
	if err != nil {
		return nil, err
	}

	res.ProblemID = p.ID
	if p.Answer == answer.Answer || p.Type == 3 {
		res.Result = true
	} else {
		res.Result = false
		res.Answer = p.Answer
	}

	return res, nil
}
