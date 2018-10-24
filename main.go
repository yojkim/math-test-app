package main

import (
	"github.com/yojkim/math-test-app/internal/infrastructures"
	"github.com/yojkim/math-test-app/internal/infrastructures/sqlite"
)

func main() {
	// disconnect db connection when process is ended.
	defer sqlite.Close()

	infrastructures.Router.Start(":3000")
}
