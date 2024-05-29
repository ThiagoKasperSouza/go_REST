package repositoryscenario

import (
	r_steps "newsRestFiber/test/scenarios/repository_scenario/steps"

	"github.com/cucumber/godog"
)

func myRedisClientIsRunning() error {
	return r_steps.ClientExists()
}

func iCreateANewUserWithTheFollowingDetails(arg1 *godog.Table) error {
	return r_steps.CreateObj()

}

func iShouldBeAbleToGetItByKey() error {
	return r_steps.GetObj()
}
func iShouldBeAbleToListObjectsFromTheExampleKey() error {
	return r_steps.GetAll()
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my redis client is running$`, myRedisClientIsRunning)
	ctx.Step(`^i create a new user with the following details:$`, iCreateANewUserWithTheFollowingDetails)
	ctx.Step(`^i should be able to get it by key$`, iShouldBeAbleToGetItByKey)
	ctx.Step(`^i should be able to list objects from the example key$`, iShouldBeAbleToListObjectsFromTheExampleKey)
}
