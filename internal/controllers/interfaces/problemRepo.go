package interfaces

import "github.com/yojkim/math-test-app/internal/domains"

type ProblemRepository interface {
	Store(problem domains.Problem) (int, error)
	FindAll() ([]domains.Problem, error)
	FindById(id int) (*domains.Problem, error)
}
