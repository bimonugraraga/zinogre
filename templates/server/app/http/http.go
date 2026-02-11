package app

import (
	resource "github.com/bimonugraraga/zinogre/templates/server/dal/resources"
	"github.com/bimonugraraga/zinogre/templates/server/routes"
)

func Main() {
	resource.InitNewRepo()
	resource.InitNewUsecase()
	routes.InitFiber()
}
