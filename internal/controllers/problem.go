package controllers

import (
	"github.com/yojkim/math-test-app/internal/controllers/interfaces"
	"github.com/yojkim/math-test-app/internal/domains"
)

type ProblemController struct {
	Repository interfaces.ProblemRepository
}

func (ctl *ProblemController) AddProblems(problems []domains.Problem) error {
	for _, p := range problems {
		_, err := ctl.Repository.Store(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctl *ProblemController) GetAllProblems() ([]domains.Problem, error) {
	return ctl.Repository.FindAll()
}
