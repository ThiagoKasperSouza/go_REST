// redis_test.go
package main

import (
	"testing"

	r_steps "newsRestFiber/test/steps"

	"github.com/cucumber/godog"
	"github.com/redis/go-redis/v9"
)

// godogsCtxKey is the key used to store the available godogs in the context.Context.
type defaultCtx struct {
	Client *redis.Client
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.Step(`[Repo] I create an example Obj ->`, r_steps.CreateObj)
}
