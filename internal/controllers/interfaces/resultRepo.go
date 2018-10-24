package interfaces

import "github.com/yojkim/math-test-app/internal/domains"

type ResultRepository interface {
	Store(domains.Result) (int, error)
	Check(*domains.Result) (*domains.Result, error)
}
