package services

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/yojkim/math-test-app/internal/controllers"
	"github.com/yojkim/math-test-app/internal/domains"
	"github.com/yojkim/math-test-app/internal/repositories"
	"github.com/yojkim/math-test-app/internal/services/interfaces"
	"github.com/yojkim/math-test-app/internal/viewmodels"
	"log"
	"net/http"
)

type ProblemService struct {
	Controller controllers.ProblemController
}

func NewProblemService(conn *gorm.DB) *ProblemService {
	return &ProblemService{
		Controller: controllers.ProblemController{
			Repository: &repositories.ProblemRepository{
				Conn: conn,
			},
		},
	}
}

func (service *ProblemService) GetAllProblems(c interfaces.Context) error {
	problems, err := service.Controller.GetAllProblems()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	outputDatas := make([]viewmodels.Problem, 0)
	for _, problem := range problems {
		data := viewmodels.Problem{
			ID:      problem.ID,
			Text:    problem.Text,
			Type:    problem.Type,
			Choices: problem.Choices,
		}

		outputDatas = append(outputDatas, data)
	}

	return c.JSON(http.StatusOK, viewmodels.ProblemsVM{Problems: outputDatas})
}

func (service *ProblemService) CreateProblems(c interfaces.Context) error {
	var problems []domains.Problem
	inputData := c.FormValue("problems")

	err := json.Unmarshal([]byte(inputData), &problems)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = service.Controller.AddProblems(problems)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return c.String(http.StatusCreated, "")
}
