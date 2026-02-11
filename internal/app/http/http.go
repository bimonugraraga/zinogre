package app

import (
	resource "github.com/bimonugraraga/zinogre/internal/dal/resources"
	"github.com/bimonugraraga/zinogre/internal/routes"
)

func Main() {
	resource.InitNewRepo()
	resource.InitNewUsecase()
	routes.InitFiber()
}
