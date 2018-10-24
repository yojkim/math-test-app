package controllers

import (
	"github.com/yojkim/math-test-app/internal/controllers/interfaces"
	"github.com/yojkim/math-test-app/internal/domains"
)

type ResultController struct {
	Repository interfaces.ResultRepository
}

func (ctl *ResultController) CheckAnswer(answers []domains.Answer) ([]domains.Result, error) {
	results := make([]domains.Result, 0)
	for _, answer := range answers {
		result, err := ctl.Repository.Check(&answer)
		if err != nil {
			return nil, err
		}

		_, err = ctl.Repository.Store(*result)
		if err != nil {
			return nil, err
		}

		results = append(results, *result)
	}

	return results, nil
}
