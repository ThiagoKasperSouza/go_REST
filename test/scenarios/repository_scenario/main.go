package repositoryscenario

import (
	r_steps "newsRestFiber/test/scenarios/repository_scenario/steps"

	"github.com/cucumber/godog"
)

func iCreateANewUserWithTheFollowingDetails(arg1 *godog.Table) error {
	return r_steps.CreateObj()

}

func iShouldBeAbleToGetItByKey() error {
	return r_steps.GetObj()
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^i create a new user with the following details:$`, iCreateANewUserWithTheFollowingDetails)
	ctx.Step(`^i should be able to get it by key$`, iShouldBeAbleToGetItByKey)
}
