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

func (repo *ResultRepository) Check(answer *domains.Result) (*domains.Result, error) {
	p := domains.Problem{}

	err := repo.Conn.First(&p).Error
	if err != nil {
		return nil, err
	}

	// todo : need to check picture case
	if p.Answer == answer.Answer {
		answer.Result = 1
	} else {
		answer.Result = 0
		answer.Answer = p.Answer
	}

	return answer, nil
}
