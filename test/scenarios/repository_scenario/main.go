package repositoryscenario

import (
	r_steps "go_Rest/test/scenarios/repository_scenario/steps"

	"github.com/cucumber/godog"
)

/*
Copyright 2024 Thiago Kasper de Souza

This file is part of rsNews_blogApi.

rsNews_blogApi is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

rsNews_blogApi is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with rsNews_blogApi.  If not, see <https://www.gnu.org/licenses/>
*/

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
