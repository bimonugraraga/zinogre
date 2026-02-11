package app

import (
	resource "github.com/bimonugraraga/zinogre/generator/templates/server/dal/resources"
	"github.com/bimonugraraga/zinogre/generator/templates/server/routes"
)

func Main() {
	resource.InitNewRepo()
	resource.InitNewUsecase()
	routes.InitFiber()
}
