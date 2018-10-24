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

type ResultService struct {
	Controller controllers.ResultController
}

func NewResultService(conn *gorm.DB) *ResultService {
	return &ResultService{
		Controller: controllers.ResultController{
			Repository: &repositories.ResultRepository{
				Conn: conn,
			},
		},
	}
}

func (service *ResultService) CheckAnswer(c interfaces.Context) error {
	var answers []domains.Answer
	inputData := c.FormValue("input")

	err := json.Unmarshal([]byte(inputData), &answers)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	results, err := service.Controller.CheckAnswer(answers)
	if err != nil {
		return err
	}

	outputDatas := make([]viewmodels.Result, 0)
	for _, result := range results {
		data := viewmodels.Result{
			ProblemID: result.ProblemID,
			Result:    result.Result,
			Answer:    result.Answer,
		}

		outputDatas = append(outputDatas, data)
	}

	return c.JSON(http.StatusOK, viewmodels.ResultsVM{Results: outputDatas})
}
