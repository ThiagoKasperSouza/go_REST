// redis_test.go
package main

import (
	"testing"

	r "newsRestFiber/repository"
	r_scen "newsRestFiber/test/scenarios/repository_scenario"

	"github.com/cucumber/godog"
)

var rdb = r.DbClient{
	Instance: r.GetClient(),
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"./test/features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	r_scen.InitializeScenario(ctx)
}
