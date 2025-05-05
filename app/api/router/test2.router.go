package router

import (
	"workflower/app/core"
	"workflower/app/lib"
	"workflower/app/service"
)

type Test2Routes struct {
	logger      lib.Logger
	engine      core.Engine
	testService service.TestService
}

func NewTest2Routes(
	logger lib.Logger,
	engine core.Engine,
	testService service.TestService,

) Test2Routes {
	return Test2Routes{
		engine:      engine,
		logger:      logger,
		testService: testService,
	}
}

func (t Test2Routes) Setup() {
	testRoutes := t.engine.ApiGroup.Group("/test")
	{
		testRoutes.GET("", t.testService.Test)

	}
}
