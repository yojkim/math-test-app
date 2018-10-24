package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yojkim/math-test-app/internal/domains"
)

type ProblemRepository struct {
	Conn *gorm.DB
}

func (repo *ProblemRepository) Store(problem domains.Problem) (int, error) {
	err := repo.Conn.Create(&problem).Error
	if err != nil {
		return 0, err
	}

	return problem.ID, nil
}

func (repo *ProblemRepository) FindAll() ([]domains.Problem, error) {
	var p []domains.Problem

	err := repo.Conn.Find(&p).Error
	if err != nil {
		return nil, err
	}

	//n := len(p)
	//problems := make([]domains.Problem, n)
	//
	//for idx, p := range p {
	//	problems[idx] = p.bindToDomainData()
	//}

	return p, nil
}

func (repo *ProblemRepository) FindById(id int) (*domains.Problem, error) {
	p := domains.Problem{}

	err := repo.Conn.First(p, id).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

//func (p Problem) bindToDomainData() domains.Problem {
//	problem := domains.Problem{
//		ID:      int(p.Model.ID),
//		Text:    p.Text,
//		Type:    p.Type,
//		Choices: p.Choices,
//		Answer:  p.Choices,
//	}
//	return problem
//}
