package repositoryscenario

import (
	r_steps "newsRestFiber/test/scenarios/repository_scenario/steps"

	"github.com/cucumber/godog"
)

func myRedisClientIsRunning() error {
	return r_steps.ClientExists()
}

func iCreateANewObjectWithTheFollowingDetails(arg1 *godog.Table) error {
	return r_steps.CreateObj()

}

func iShouldBeAbleToGetItByKey() error {
	return r_steps.GetObj()
}
func iShouldBeAbleToListObjectsFromTheExampleKey() error {
	return r_steps.GetAll()
}

func iHaveAtLeastOneObjectOnTheClient() error {
	return r_steps.GetObj()
}

func iShouldBeAbleToDeleteThisObject() error {
	return r_steps.Delete()
}

func iShouldBeAbleToUpdateThisObject() error {
	return r_steps.Update()
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my redis client is running$`, myRedisClientIsRunning)
	ctx.Step(`^i create a new object with the following details:$`, iCreateANewObjectWithTheFollowingDetails)
	ctx.Step(`^i should be able to get it by key$`, iShouldBeAbleToGetItByKey)
	ctx.Step(`^i should be able to list objects from the example key$`, iShouldBeAbleToListObjectsFromTheExampleKey)
	ctx.Step(`^i have at least one object on the client$`, iHaveAtLeastOneObjectOnTheClient)
	ctx.Step(`^i should be able to delete this object$`, iShouldBeAbleToDeleteThisObject)
	ctx.Step(`^i should be able to update this object$`, iShouldBeAbleToUpdateThisObject)
}
