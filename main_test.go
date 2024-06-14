package main

import (
	"testing"

	r "go_Rest/src/repository"
	r_scen "go_Rest/test/scenarios/repository_scenario"

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
